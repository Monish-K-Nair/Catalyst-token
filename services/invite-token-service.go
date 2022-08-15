package services

import (
	repository "catalyst-token/controllers/invite-token-controllers"
	model "catalyst-token/models/invite-token-models"
)

type invite_service struct {
	repo repository.Repository
}

func InviteServiceRegister(repo repository.Repository) *invite_service {
	return &invite_service{repo: repo}
}

type InviteService interface {
	ValidateToken(input string) bool
	ListToken() (*[]model.InviteToken)
	GenerateToken() (string, string)
	RevokeToken(input string) (bool, string)
	DeleteToken(input string) (bool, string)
}

func (s *invite_service) ValidateToken(input string) bool {
	val := s.repo.ValidateToken(input)
	return val
}

func (s *invite_service) ListToken() (*[]model.InviteToken) {
	val := s.repo.RetrieveTokens()
	return val
}

func (s *invite_service) GenerateToken() (string, string) {
	val, status := s.repo.GenerateToken()
	return val, status
}

func (s *invite_service) RevokeToken(input string) (bool, string) {

	val, err := s.repo.UpdateToken(input)
	return val, err
}

func (s *invite_service) DeleteToken(input string) (bool, string) {

	val, err := s.repo.DeleteToken(input)
	return val, err
}
