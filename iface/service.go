package mdgeotrack

import (
	"context"
	emailverification "md-auth-svc/request_response/email_verification"
	signin "md-auth-svc/request_response/sign_in"
	signup "md-auth-svc/request_response/sign_up"
)

type Service interface {
	// Health check
	HeartBeat() map[string]string
	// User API's
	UserSignUp(ctx context.Context, req signup.UserSignUpRequest) (*signup.SignUpResponse, error)
	EmailVerification(ctx context.Context, req emailverification.VerifyEmailRequest) (*emailverification.VerifyEmailResponse, error)
	SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error)
	// SysAdmin API's
	SysAdminSignUp(ctx context.Context, req signup.SysAdminSignUpRequest) (*signup.SignUpResponse, error)
}
