package cognito

import (
	"context"
	"fmt"
	resetpassword "md-auth-svc/request_response/reset_password"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// ResetPassword completes the password reset process for a user.
// It takes a context and a ResetPasswordRequest, and returns a ResetPasswordResponse or an error.
func (s *AuthService) ResetPassword(ctx context.Context, req *resetpassword.ResetPasswordRequest) (*resetpassword.ResetPasswordResponse, error) {
	input := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         &s.cognito.ClientID,
		Username:         &req.Username,
		ConfirmationCode: &req.ConfirmationCode,
		Password:         &req.NewPassword,
	}

	_, err := s.cognito.Client.ConfirmForgotPassword(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("password reset failed: %v", err)
	}

	return &resetpassword.ResetPasswordResponse{
		Message: "Password reset successful",
	}, nil
}
