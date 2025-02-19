package repository

import (
	"context"
	"fmt"
	signup "md-auth-svc/request_response/sign_up"
	sysadminuser "md-auth-svc/request_response/sys_admin_user"
	userinfo "md-auth-svc/request_response/tenant_user_info"
	"time"

	"github.com/google/uuid"
)

// UserSignUp
func (r *repository) UserSignUp(ctx context.Context, req *signup.UserSignUpRequest, uniqueID string) (*signup.SignUpResponse, error) {

	uniqueUUID, err := uuid.Parse(uniqueID)
	if err != nil {
		return nil, fmt.Errorf("invalid uniqueID: %v", err)
	}

	tenantUUID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("invalid TenantID: %v", err)
	}

	user := &userinfo.User{
		ID:         uniqueUUID, // Cognito user ID
		Username:   req.Username,
		Email:      req.Email,
		GivenName:  req.GivenName,
		FamilyName: req.FamilyName,
		Role:       req.UserPoolGroup,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		TenantID:   tenantUUID,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to save user in database: %v", err)
	}

	return &signup.SignUpResponse{
		UserSub: uniqueID,
		Message: "User signed up successfully!",
	}, nil
}

// SysAdminSignUp registers a new system admin in the database.
func (r *repository) SysAdminSignUp(ctx context.Context, req *signup.SysAdminSignUpRequest, uniqueID string) (*signup.SignUpResponse, error) {
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
