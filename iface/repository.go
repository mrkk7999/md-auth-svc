package mdgeotrack

import (
	"context"
	signup "md-auth-svc/request_response/sign_up"
	sysadminuser "md-auth-svc/request_response/sys_admin_user"
	userinfo "md-auth-svc/request_response/tenant_user_info"
)

type Repository interface {
	// Health check
	HeartBeat() map[string]string

	// System Admin API's
	CreateSystemAdmin(ctx context.Context, req *signup.SysAdminSignUpRequest, uniqueID string) (*signup.SignUpResponse, error)
	GetAllSysAdmins(ctx context.Context) ([]sysadminuser.SysAdmin, error)
	GetSysAdminByID(ctx context.Context, adminID string) (*sysadminuser.SysAdmin, error)
	// UpdateSysAdmin(ctx context.Context, admin *sysadminuser.SysAdmin) error
	// DeleteSysAdmin(ctx context.Context, username, adminID string) error

	// Tenant User API's
	CreateTenantUser(ctx context.Context, req *signup.UserSignUpRequest, uniqueID string) (*signup.SignUpResponse, error)
	GetAllTenantUsers(ctx context.Context) ([]userinfo.User, error)
	GetTenantUserByID(ctx context.Context, userID string) (*userinfo.User, error)
	GetUsersByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error)
	GetAdminByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error)
}
