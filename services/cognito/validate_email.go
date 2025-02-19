package cognito

import (
	"context"
	"fmt"
	emailverification "md-auth-svc/request_response/email_verification"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// VerifyEmail - Confirms a user's email using the verification code from Cognito
func (s *AuthService) VerifyEmail(ctx context.Context, req *emailverification.VerifyEmailRequest) (*emailverification.VerifyEmailResponse, error) {
	secretHash := s.cognito.GenerateSecretHash(req.Username)

	// Step 1:
	// Confirm user sign-up with verification code
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         &s.cognito.ClientID,
		Username:         &req.Username,
		ConfirmationCode: &req.VerificationCode,
		SecretHash:       &secretHash,
	}

	_, err := s.cognito.Client.ConfirmSignUp(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("email verification failed: %v", err)
	}

	return &emailverification.VerifyEmailResponse{
		Message: "Email verified successfully",
	}, nil
}
