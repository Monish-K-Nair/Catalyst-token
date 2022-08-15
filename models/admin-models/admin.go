package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	utils "catalyst-token/utils"
)

type Admin struct {
	ID        string `gorm:"primaryKey;"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (admin_user *Admin) BeforeCreate(db *gorm.DB) error {
	admin_user.ID = uuid.New().String()
	admin_user.Password = utils.HashPassword(admin_user.Password)
	admin_user.CreatedAt = time.Now().Local()
	return nil
}

func (admin_user *Admin) BeforeUpdate(db *gorm.DB) error {
	admin_user.UpdatedAt = time.Now().Local()
	return nil
}
