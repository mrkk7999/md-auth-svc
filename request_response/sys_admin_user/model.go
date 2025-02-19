package sysadminuser

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SysAdmin struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username   string    `gorm:"unique;not null"`
	Email      string    `gorm:"unique;not null"`
	GivenName  string    `gorm:"not null"`
	FamilyName string    `gorm:"not null"`
	// PasswordHash string     `gorm:"not null"`
	IsDeleted bool       `gorm:"default:false"`
	DeletedAt *time.Time `gorm:"default:NULL"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

// Enable soft delete
func (admin *SysAdmin) BeforeDelete(tx *gorm.DB) (err error) {
	now := time.Now()
	admin.DeletedAt = &now
	admin.IsDeleted = true
	return nil
}

// overrides the default table name in GORM
func (SysAdmin) TableName() string {
	return "sys_admin"
}
