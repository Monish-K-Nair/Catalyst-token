package InviteTokenController

import (
	models "catalyst-token/models/invite-token-models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ValidateToken(input string) bool
	RetrieveTokens() *[]models.InviteToken
	GenerateToken() (string, bool)
	UpdateToken(input string) bool
	DeleteToken(input string) bool
}

type inv_repository struct {
	db *gorm.DB
}

func InvRepositoryRegister(db *gorm.DB) *inv_repository {
	return &inv_repository{db: db}
}

func (r *inv_repository) ValidateToken(input string) bool {

	var token models.InviteToken

	var exists bool
	err := r.db.Model(&token).
		Select("count(*) > 0").
		Where("Active = ? AND token >= ?", true, input).
		Find(&exists).
		Error

	return err == nil
}

func (r *inv_repository) RetrieveTokens() *[]models.InviteToken {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	results := db.Debug().Select("*").Where("Active = ?", true).Find(&tokens)

	if results.Error != nil {
		return &tokens
	}
	return &tokens
}

func (r *inv_repository) GenerateToken() (string, bool) {

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

func (r *inv_repository) UpdateToken(token string) bool {

	var tokens models.InviteToken
	db := r.db.Model(&tokens)

	revoke := db.Debug().Where("token = ?", token).Update("Active", false)

	return revoke.Error == nil
}

func (r *inv_repository) DeleteToken(token string) bool {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	revoke := db.Debug().Where("token = ?", token).Delete(&tokens)

	return revoke.Error == nil 
}
