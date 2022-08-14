package AdminAuth

import (
	models "catalyst-token/models"
	utils "catalyst-token/utils"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterToken(input *models.Admin) (*models.Admin, string)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) RegisterToken(input *models.Admin) (*models.Admin, string) {

	errorCode := make(chan string, 1)
	var users models.Admin
	db := r.db.Model(&users)

	admin_info := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if admin_info.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := utils.ValidatePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}


