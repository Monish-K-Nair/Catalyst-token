package invitetokenmodel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InviteToken struct {
	ID        string `gorm:"primaryKey;"`
	Token     string `gorm:"type:varchar(255);unique"`
	Name      string `gorm:"type:varchar(255);"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (tk *InviteToken) BeforeCreate(db *gorm.DB) error {
	tk.ID = uuid.New().String()
	tk.CreatedAt = time.Now().Local()
	return nil
}

func (tk *InviteToken) BeforeUpdate(db *gorm.DB) error {
	tk.UpdatedAt = time.Now().Local()
	return nil
}
