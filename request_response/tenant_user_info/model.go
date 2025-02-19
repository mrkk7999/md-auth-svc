package userinfo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TenantID   uuid.UUID `gorm:"not null;index"`
	Username   string    `gorm:"unique;not null"`
	Email      string    `gorm:"unique;not null"`
	GivenName  string
	FamilyName string
	// Role-based Access Control (RBAC)
	Role      string     `gorm:"type:varchar(50);not null;check:role IN ('admin', 'user')"`
	IsDeleted bool       `gorm:"default:false"`
	DeletedAt *time.Time `gorm:"default:NULL"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

// Enable soft delete
func (user *User) BeforeDelete(tx *gorm.DB) (err error) {
	now := time.Now()
	user.DeletedAt = &now
	user.IsDeleted = true
	return nil
}
