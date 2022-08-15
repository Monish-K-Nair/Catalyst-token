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
	ValidateToken(input *model.InviteToken) bool
	ListToken() (*[]model.InviteToken, string)
	GenerateToken() (string, bool)
	RevokeToken(input *model.InviteToken) bool
	DeleteToken(input *model.InviteToken) bool
}

func (s *invite_service) ValidateToken(input *model.InviteToken) bool {

	token := model.InviteToken{
		Token: input.Token,
	}

	val := s.repo.ValidateToken(&token)
	return val
}

func (s *invite_service) ListToken() (*[]model.InviteToken, string) {
	val := s.repo.RetrieveTokens()
	return val, ""
}

func (s *invite_service) GenerateToken() (string, bool) {
	val, status := s.repo.GenerateToken()
	return val, status
}

func (s *invite_service) RevokeToken(input *model.InviteToken) bool {

	val := s.repo.UpdateToken(input.Token)
	return val
}

func (s *invite_service) DeleteToken(input *model.InviteToken) bool {

	val := s.repo.DeleteToken(input.Token)
	return val
}
