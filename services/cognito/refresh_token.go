package cognito

import (
	"context"
	"fmt"
	refreshtoken "md-auth-svc/request_response/refresh_token"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// RefreshToken refreshes the authentication token for a user.
// It takes a context and a RefreshTokenRequest, and returns a RefreshTokenResponse or an error.
func (s *AuthService) RefreshToken(ctx context.Context, req *refreshtoken.RefreshTokenRequest) (*refreshtoken.RefreshTokenResponse, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "REFRESH_TOKEN_AUTH",
		ClientId: &s.cognito.ClientID,
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": req.RefreshToken,
		},
	}

	result, err := s.cognito.Client.InitiateAuth(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("token refresh failed: %v", err)
	}

	if result.AuthenticationResult == nil {
		return nil, fmt.Errorf("authentication result is nil")
	}

	return &refreshtoken.RefreshTokenResponse{
		AccessToken: *result.AuthenticationResult.AccessToken,
		IdToken:     *result.AuthenticationResult.IdToken,
		ExpiresIn:   int(result.AuthenticationResult.ExpiresIn),
		Message:     "Token refreshed successfully",
	}, nil
}
