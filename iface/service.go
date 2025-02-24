package mdgeotrack

import (
	"context"
	emailverification "md-auth-svc/request_response/email_verification"
	signin "md-auth-svc/request_response/sign_in"
	signup "md-auth-svc/request_response/sign_up"
	sysadminuser "md-auth-svc/request_response/sys_admin_user"
	userinfo "md-auth-svc/request_response/tenant_user_info"
)

type Service interface {
	// Health check
	HeartBeat() map[string]string

	// Common API's
	SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error)
	SignOut(ctx context.Context, accessToken string) error
	EmailVerification(ctx context.Context, req emailverification.VerifyEmailRequest) (*emailverification.VerifyEmailResponse, error)
	// ChangeGroup(ctx context.Context, req usergroup.UserToGroupRequest) error

	// SysAdmin API's
	SysAdminSignUp(ctx context.Context, req signup.SysAdminSignUpRequest) (*signup.SignUpResponse, error)
	GetAllSysAdmins(ctx context.Context) ([]sysadminuser.SysAdmin, error)
	GetSysAdminByID(ctx context.Context, adminID string) (*sysadminuser.SysAdmin, error)
	// UpdateSysAdmin(ctx context.Context, admin *sysadminuser.SysAdmin) error
	// DeleteSysAdmin(ctx context.Context, username, adminID string) error

	// User API's
	UserSignUp(ctx context.Context, req signup.UserSignUpRequest) (*signup.SignUpResponse, error)
	GetAllTenantUsers(ctx context.Context) ([]userinfo.User, error)
	GetTenantUserByID(ctx context.Context, userID string) (*userinfo.User, error)
	GetUsersByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error)
	GetAdminByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error)
}
