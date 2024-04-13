package service

import (
	"strconv"
	cache "tools/internals/cache/middleware"
	"tools/internals/models"
	"tools/internals/repository"
)

type Service struct {
	rp    *repository.Repository
	cache *cache.Cache
}

func NewService(rp *repository.Repository, cache *cache.Cache) *Service {
	return &Service{rp: rp, cache: cache}
}

func (s *Service) FindBanner(input *models.BannerGetRequest) (*models.Banner, error) {
	repoBann, err := s.rp.FindBanner(input)
	if err != nil {
		return nil, err
	}
	if repoBann.IsActive {

	}
	if !input.UseLastRevision {
		key := strconv.FormatUint(uint64(input.TagID), 10) + strconv.FormatUint(uint64(input.FeatureID), 10)
		banner, err := s.cache.FindBanner(key)
		if err != nil && err.Error() == "redis returned nil" {
			s.cache.AddFromRepo(repoBann, input)
			return repoBann, err
		}
		return banner, nil
	}
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

func (s *Service) GetFilteredBanner(input *models.BannerGetAdminRequest) ([]models.Banner, error) {
	return s.rp.GetFilteredBanner(input)
}
