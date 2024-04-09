package service

import (
	"tools/internals/models"
	"tools/internals/repository"
)

type Service struct {
	rp *repository.Repository
}

func NewService(rp *repository.Repository) *Service {
	return &Service{rp: rp}
}

func (s *Service) FindBanner(input *models.BannerGetMethod) (*models.Banner, error) {
	return s.rp.FindBanner(input)
}

func (s *Service) CreateBanner(input *models.Banner) (*models.InsertedBannerResponse, error) {
	return s.rp.CreateBanner(input)
}
