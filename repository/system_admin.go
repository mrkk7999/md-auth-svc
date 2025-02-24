package repository

import (
	"context"
	"errors"
	"fmt"
	signup "md-auth-svc/request_response/sign_up"
	sysadminuser "md-auth-svc/request_response/sys_admin_user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateSystemAdmin registers a new system admin in the database.
func (r *repository) CreateSystemAdmin(ctx context.Context, req *signup.SysAdminSignUpRequest, uniqueID string) (*signup.SignUpResponse, error) {
	uniqueUUID, err := uuid.Parse(uniqueID)
	if err != nil {
		return nil, fmt.Errorf("invalid uniqueID: %v", err)
	}

	sysAdmin := &sysadminuser.SysAdmin{
		ID:         uniqueUUID, // Cognito user ID
		Username:   req.Username,
		Email:      req.Email,
		GivenName:  req.GivenName,
		FamilyName: req.FamilyName,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save to database
	if err := r.db.Create(sysAdmin).Error; err != nil {
		return nil, fmt.Errorf("failed to save sys_admin in database: %v", err)
	}

	return &signup.SignUpResponse{
		UserSub: uniqueID,
		Message: "SysAdmin signed up successfully!",
	}, nil
}

// GetAllSysAdmins retrieves all system admins from the database.
func (r *repository) GetAllSysAdmins(ctx context.Context) ([]sysadminuser.SysAdmin, error) {
	var admins []sysadminuser.SysAdmin
	err := r.db.WithContext(ctx).Find(&admins).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch system admins: %v", err)
	}
	return admins, nil
}

// GetSysAdminByID retrieves a system admin by their ID.
func (r *repository) GetSysAdminByID(ctx context.Context, adminID string) (*sysadminuser.SysAdmin, error) {
	var admin sysadminuser.SysAdmin
	err := r.db.WithContext(ctx).Where("id = ?", adminID).First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch system admin: %v", err)
	}
	return &admin, nil
}

// // UpdateSysAdmin updates a system admin's details, prioritizing username over ID.
// func (r *repository) UpdateSysAdmin(ctx context.Context, admin *sysadminuser.SysAdmin) error {
// 	query := r.db.WithContext(ctx).Model(&sysadminuser.SysAdmin{})

// 	// Prioritize ID , then fallback to username
// 	if admin.ID != uuid.Nil {
// 		query = query.Where("id = ?", admin.ID)

// 	} else {
// 		return errors.New("user ID must be provided for update")
// 	}

// 	// Create an update map with non-empty fields
// 	updateFields := map[string]interface{}{}

// 	if admin.GivenName != "" {
// 		updateFields["given_name"] = admin.GivenName
// 	}
// 	if admin.FamilyName != "" {
// 		updateFields["family_name"] = admin.FamilyName
// 	}

// 	// Prevent updating ID and Email
// 	if len(updateFields) == 0 {
// 		return errors.New("no valid fields provided for update")
// 	}

// 	updateFields["updated_at"] = time.Now()

// 	result := query.Updates(updateFields)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("system admin not found or no changes detected")
// 	}
// 	return nil
// }

// // DeleteSysAdmin deletes a system admin, prioritizing username first, then ID.
// func (r *repository) DeleteSysAdmin(ctx context.Context, username, adminID string) error {
// 	query := r.db.WithContext(ctx).Model(&sysadminuser.SysAdmin{})

// 	// Prioritize username, then fallback to ID
// 	if adminID != "" {
// 		query = query.Where("id = ?", adminID)
// 	} else {
// 		return errors.New("user ID must be provided for deletion")
// 	}

// 	result := query.Delete(&sysadminuser.SysAdmin{})
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("system admin not found")
// 	}
// 	return nil
// }
