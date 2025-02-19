package mdgeotrack

import (
	"context"
	signup "md-auth-svc/request_response/sign_up"
)

type Repository interface {
	// Health check
	HeartBeat() map[string]string

	//
	UserSignUp(ctx context.Context, req *signup.UserSignUpRequest, uniqueID string) (*signup.SignUpResponse, error)
	SysAdminSignUp(ctx context.Context, req *signup.SysAdminSignUpRequest, uniqueID string) (*signup.SignUpResponse, error)
}
