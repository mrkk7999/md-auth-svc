package repository

import (
	mdgeotrack "md-auth-svc/iface"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) mdgeotrack.Repository {
	return &repository{
		db: db,
	}
}
