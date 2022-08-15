package InviteTokenController

import (
	models "catalyst-token/models/invite-token-models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ValidateToken(input *models.InviteToken) bool
	RetrieveTokens() *[]models.InviteToken
	GenerateToken() (string, bool)
	UpdateToken(string) bool
	DeleteToken(string) bool
}

type repository struct {
	db *gorm.DB
}

func (r *repository) ValidateToken(input *models.InviteToken) bool {

	var token models.InviteToken

	var exists bool
	err := r.db.Model(&token).
		Select("count(*) > 0").
		Where("Active = ? AND token >= ?", true, input.Token).
		Find(&exists).
		Error

	return err == nil
}

func (r *repository) RetrieveTokens() *[]models.InviteToken {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	results := db.Debug().Select("*").Where("Active = ?", true).Find(&tokens)

	if results.Error != nil {
		return &tokens
	}
	return &tokens
}

func (r *repository) GenerateToken() (string, bool) {

	var tokens models.InviteToken
	db := r.db.Model(&tokens)

	tokens.Active = true
	tokens.Token = strings.Replace(uuid.New().String(), "-", "", -1)[:12]
	tokens.Name = ""

	insert := db.Debug().Create(&tokens)
	db.Commit()


	if insert.Error != nil {
		return "", false
	}
	return tokens.Token, true
}

func (r *repository) UpdateToken(token string) bool {

	var tokens models.InviteToken
	db := r.db.Model(&tokens)

	revoke := db.Debug().Where("token = ?", token).Update("Active", false)

	return revoke.Error == nil
}

func (r *repository) DeleteToken(token string) bool {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	revoke := db.Debug().Where("token = ?", token).Delete(&tokens)

	return revoke.Error == nil 
}
