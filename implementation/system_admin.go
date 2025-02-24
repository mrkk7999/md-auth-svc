package implementation

import (
	"context"
	sysadminuser "md-auth-svc/request_response/sys_admin_user"
)

// GetAllSysAdmins retrieves all system admins
func (s *service) GetAllSysAdmins(ctx context.Context) ([]sysadminuser.SysAdmin, error) {
	return s.repository.GetAllSysAdmins(ctx)
}

// GetSysAdminByID fetches a system admin by their ID
func (s *service) GetSysAdminByID(ctx context.Context, adminID string) (*sysadminuser.SysAdmin, error) {
	return s.repository.GetSysAdminByID(ctx, adminID)
}

// // UpdateSysAdmin updates a system admin, ensuring only non-null fields are modified
// func (s *service) UpdateSysAdmin(ctx context.Context, admin *sysadminuser.SysAdmin) error {
// 	return s.repository.UpdateSysAdmin(ctx, admin)
// }

// // DeleteSysAdmin handles soft deletion of a system admin
// func (s *service) DeleteSysAdmin(ctx context.Context, username, adminID string) error {
// 	return s.repository.DeleteSysAdmin(ctx, username, adminID)
// }
