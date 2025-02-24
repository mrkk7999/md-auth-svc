package implementation

import (
	"context"
	userinfo "md-auth-svc/request_response/tenant_user_info"
)

func (s *service) GetAllTenantUsers(ctx context.Context) ([]userinfo.User, error) {
	return s.repository.GetAllTenantUsers(ctx)
}

func (s *service) GetTenantUserByID(ctx context.Context, userID string) (*userinfo.User, error) {
	return s.repository.GetTenantUserByID(ctx, userID)
}

func (s *service) GetUsersByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error) {
	return s.repository.GetUsersByTenantID(ctx, tenantID)
}

func (s *service) GetAdminByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error) {
	return s.repository.GetAdminByTenantID(ctx, tenantID)
}
