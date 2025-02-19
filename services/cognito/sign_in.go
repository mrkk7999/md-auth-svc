package cognito

import (
	"context"
	"fmt"
	"log"
	signin "md-auth-svc/request_response/sign_in"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/golang-jwt/jwt/v4"
)

// SignIn authenticates a user with their credentials.
// It takes a context and a SignInRequest, and returns a SignInResponse or an error.
func (s *AuthService) SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error) {
	// Generate SecretHash for the request
	secretHash := s.cognito.GenerateSecretHash(req.Username)

	// Step 1: Initiate authentication
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: &s.cognito.ClientID,
		AuthParameters: map[string]string{
			"USERNAME":    req.Username,
			"PASSWORD":    req.Password,
			"SECRET_HASH": secretHash, // ✅ Include SecretHash
		},
	}

	authResp, err := s.cognito.Client.InitiateAuth(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}
	// Step 2: Decode JWT to check expiration time
	idToken := *authResp.AuthenticationResult.IdToken
	expirationTime, err := getTokenExpiration(idToken)
	if err != nil {
		log.Println("Error decoding token:", err)
	} else {
		log.Println("ID Token Expiration Time:", expirationTime)
	}
	// Step 2: Return JWT tokens
	return &signin.SignInResponse{
		AccessToken:  *authResp.AuthenticationResult.AccessToken,
		RefreshToken: *authResp.AuthenticationResult.RefreshToken,
		IdToken:      *authResp.AuthenticationResult.IdToken,
		Message:      "Login successful",
	}, nil
}

// getTokenExpiration - Extracts the expiration time from a JWT
func getTokenExpiration(tokenString string) (time.Time, error) {
	// Parse token but don't validate signature
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return time.Time{}, fmt.Errorf("invalid token claims")
	}

	// Get expiration time
	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("invalid expiration time format")
	}

	// Convert UNIX timestamp to time.Time
	expTime := time.Unix(int64(expFloat), 0)
	return expTime, nil
}

// package cognito

// import (
// 	"context"
// 	"fmt"
// 	signin "md-auth-svc/request_response/sign_in"

// 	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
// 	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
// )

// // SignIn authenticates a user
// func (s *AuthService) SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error) {
// 	input := &cognitoidentityprovider.InitiateAuthInput{
// 		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
// 		ClientId: &s.cognito.ClientID,
// 		AuthParameters: map[string]string{
// 			"USERNAME": req.Username,
// 			"PASSWORD": req.Password,
// 		},
// 	}

// 	result, err := s.cognito.Client.InitiateAuth(ctx, input)
// 	if err != nil {
// 		return nil, fmt.Errorf("sign-in failed: %v", err)
// 	}

// 	return &signin.SignInResponse{
// 		AccessToken:  *result.AuthenticationResult.AccessToken,
// 		RefreshToken: *result.AuthenticationResult.RefreshToken,
// 		IdToken:      *result.AuthenticationResult.IdToken,
// 		ExpiresIn:    int(result.AuthenticationResult.ExpiresIn),
// 		Message:      "Sign-in successful",
// 	}, nil
// }
