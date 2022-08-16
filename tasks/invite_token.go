package task

import (
	models "catalyst-token/models/invite-token-models"
	"time"

	"gorm.io/gorm"
)

func InvalidateToken(db *gorm.DB) {
	tokens := models.InviteToken{}
	db.Model(&tokens).Where("Active = ? AND Created_at <= ?", true, time.Now().UTC().AddDate(0, 0, -7)).Update("Active", false)
}
