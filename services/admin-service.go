package services

import (
	repository "catalyst-token/controllers/auth-controllers"
	model "catalyst-token/models/admin-models"
)

type InputLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Service interface {
	RegisterNewToken(input *InputLogin) (*model.Admin, error)
}

type service struct {
	repo repository.Repository
}

func ServiceRegister(repository repository.Repository) *service {
	return &service{repo: repository}
}

func (s *service) RegisterNewToken(input *InputLogin) (*model.Admin, error) {
	user := model.Admin{
		Email:    input.Email,
		Password: input.Password,
	}
	res, err := s.repo.RegisterToken(&user)
	return res, err
}
