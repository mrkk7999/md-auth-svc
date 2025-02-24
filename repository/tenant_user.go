package repository

import (
	"context"
	"errors"
	"fmt"
	signup "md-auth-svc/request_response/sign_up"
	userinfo "md-auth-svc/request_response/tenant_user_info"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateUser
func (r *repository) CreateTenantUser(ctx context.Context, req *signup.UserSignUpRequest, uniqueID string) (*signup.SignUpResponse, error) {
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
	err = r.db.Create(user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user in database: %v", err)
	}
	return &signup.SignUpResponse{
		UserSub: uniqueID,
		Message: "User signed up successfully!",
	}, nil
}

// GetAllUsers retrieves all users
func (r *repository) GetAllTenantUsers(ctx context.Context) ([]userinfo.User, error) {
	var users []userinfo.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID
func (r *repository) GetTenantUserByID(ctx context.Context, userID string) (*userinfo.User, error) {
	var user userinfo.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// UpdateUser
func (r *repository) UpdateUser(ctx context.Context, user *userinfo.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// GetUsersByTenantID
func (r *repository) GetUsersByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error) {
	var users []userinfo.User
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND role = ?", tenantID, "user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByTenantID
func (r *repository) GetAdminByTenantID(ctx context.Context, tenantID string) ([]userinfo.User, error) {
	var users []userinfo.User
	err := r.db.WithContext(ctx).Where("tenant_id = ? AND role = ?", tenantID, "admin").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser deletes a user from the database, prioritizing username first, then userID
func (r *repository) DeleteUser(ctx context.Context, username, userID string) error {
	query := r.db.WithContext(ctx).Model(&userinfo.User{})

	// Check username first, then fallback to userID
	if username != "" {
		query = query.Where("username = ?", username)
	} else if userID != "" {
		query = query.Where("id = ?", userID)
	} else {
		return errors.New("username or userID must be provided")
	}

	result := query.Delete(&userinfo.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
