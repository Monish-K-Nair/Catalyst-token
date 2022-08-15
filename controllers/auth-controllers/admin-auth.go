package AdminAuth

import (
	models "catalyst-token/models/admin-models"
	// utils "catalyst-token/utils"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterToken(input *models.Admin) (*models.Admin, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterToken(input *models.Admin) (*models.Admin, error) {

	var users models.Admin
	db := r.db.Model(&users)

	admin_info := db.Debug().Select("*").Where("email = ? AND password = ?", input.Email, input.Password).Find(&users)

	if admin_info.RowsAffected < 1 {
		return &users, errors.New("user not found")
	}

	// comparePassword := utils.ValidatePassword(&users.Password, input.Password)

	// if comparePassword != nil {
	// 	return &users, errors.New("user not found.")
	// }

	return &users, nil
}
