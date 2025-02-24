package mdgeotrack

import (
	"context"
	emailverification "md-auth-svc/request_response/email_verification"
	forgotpassword "md-auth-svc/request_response/forgot_password"
	refreshtoken "md-auth-svc/request_response/refresh_token"
	resetpassword "md-auth-svc/request_response/reset_password"
	signin "md-auth-svc/request_response/sign_in"
	signup "md-auth-svc/request_response/sign_up"
	usergroup "md-auth-svc/request_response/user_group"
	validatetoken "md-auth-svc/request_response/validate_token"
)

// CognitoServiceInterface defines the methods for interacting with the Cognito service.
// It includes functionalities for user authentication, password management, and token validation.
type CognitoServiceInterface interface {
	ForgotPassword(ctx context.Context, req *forgotpassword.ForgotPasswordRequest) (*forgotpassword.ForgotPasswordResponse, error)

	RefreshToken(ctx context.Context, req *refreshtoken.RefreshTokenRequest) (*refreshtoken.RefreshTokenResponse, error)

	ResetPassword(ctx context.Context, req *resetpassword.ResetPasswordRequest) (*resetpassword.ResetPasswordResponse, error)

	SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error)

	SignUp(ctx context.Context, req *signup.SignUpRequest) (*signup.SignUpResponse, error)

	SignOut(ctx context.Context, accessToken string) error

	ValidateToken(ctx context.Context, req *validatetoken.ValidateTokenRequest) (*validatetoken.ValidateTokenResponse, error)

	VerifyEmail(ctx context.Context, req *emailverification.VerifyEmailRequest) (*emailverification.VerifyEmailResponse, error)

	MoveUserToGroup(ctx context.Context, req usergroup.UserToGroupRequest) error
}
