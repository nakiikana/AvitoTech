package service

import (
	"fmt"
	"tools/internals/models"
	"tools/internals/repository"
)

type Service struct {
	rp *repository.Repository
}

func NewService(rp *repository.Repository) *Service {
	return &Service{rp: rp}
}

func (s *Service) FindBanner(input *models.BannerGetRequest) (*models.Banner, error) {
	return s.rp.FindBanner(input)
}

func (s *Service) CreateBanner(input *models.Banner) (*models.BannerID, error) {
	return s.rp.CreateBanner(input)
}

func (s *Service) DeleteBanner(input *models.BannerID) error {
	return s.rp.DeleteBanner(input)
}

func (s *Service) UpdateBanner(input *models.BannerUpdateRequest) error {
	return s.rp.UpdateBanner(input)
}

func (s *Service) GetFilteredBanner(input *models.BannerGetAdminRequest) (*[]models.Banner, error, int) {
	fmt.Println("here")
	return s.rp.GetFilteredBanner(input)
}
