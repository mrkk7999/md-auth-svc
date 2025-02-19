package cognito

import (
	"context"
	"fmt"
	forgotpassword "md-auth-svc/request_response/forgot_password"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// ForgotPassword initiates the password reset process for a user.
// It takes a context and a ForgotPasswordRequest, and returns a ForgotPasswordResponse or an error.
func (s *AuthService) ForgotPassword(ctx context.Context, req *forgotpassword.ForgotPasswordRequest) (*forgotpassword.ForgotPasswordResponse, error) {
	input := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: &s.cognito.ClientID,
		Username: &req.Username,
	}

	_, err := s.cognito.Client.ForgotPassword(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("forgot password request failed: %v", err)
	}

	return &forgotpassword.ForgotPasswordResponse{
		Message: "Password reset code sent",
	}, nil
}
