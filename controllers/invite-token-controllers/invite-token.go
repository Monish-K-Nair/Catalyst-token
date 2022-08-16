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
	GenerateToken() (string, string)
	UpdateToken(input string) (bool, string)
	DeleteToken(input string) (bool, string)
}

type inv_repository struct {
	db *gorm.DB
}

func InvRepositoryRegister(db *gorm.DB) *inv_repository {
	return &inv_repository{db: db}
}

func (r *inv_repository) ValidateToken(input string) bool {

	var token models.InviteToken

	err := r.db.Model(&token).Debug().
		Where("Active = ? AND Token = ?", true, input).
		Find(&token)
	return err.RowsAffected >= 1
}

func (r *inv_repository) RetrieveTokens() *[]models.InviteToken {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	results := db.Debug().Find(&tokens)

	if results.Error != nil {
		return &tokens
	}
	return &tokens
}

func (r *inv_repository) GenerateToken() (string, string) {

	var tokens models.InviteToken
	db := r.db.Model(&tokens)

	tokens.Active = true
	tokens.Token = strings.Replace(uuid.New().String(), "-", "", -1)[:12]
	tokens.Name = ""

	insert := db.Debug().Create(&tokens)
	db.Commit()

	if insert.Error != nil {
		return "", insert.Error.Error()
	}
	return tokens.Token, ""
}

func (r *inv_repository) UpdateToken(token string) (bool, string) {

	var tokens models.InviteToken
	db := r.db.Model(&tokens)

	revoke := db.Debug().Where("token = ?", token).Update("Active", false)

	if revoke.Error != nil {
		return false, revoke.Error.Error()
	}
	return true, ""
}

func (r *inv_repository) DeleteToken(token string) (bool, string) {

	var tokens []models.InviteToken
	db := r.db.Model(&tokens)

	deletion := db.Debug().Where("token = ?", token).Delete(&tokens)

	if deletion.Error != nil {
		return false, deletion.Error.Error()
	}
	return true, ""
}
