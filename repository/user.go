package repository

import (
	"context"
	"errors"
	userinfo "md-auth-svc/request_response/tenant_user_info"

	"gorm.io/gorm"
)

// CreateUser
func (r *repository) CreateUser(ctx context.Context, user *userinfo.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetUserByID
func (r *repository) GetUserByID(ctx context.Context, userID string) (*userinfo.User, error) {
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
