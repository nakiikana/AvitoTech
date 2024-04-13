package handler

import "errors"

var (
	ErrorBannerTagIdNotFound     = errors.New("GetBanner: can't get tag_id")
	ErrorBannerFeatureIdNotFound = errors.New("GetBanner: can't get feature_id")
	ErrorBannerIdFound           = errors.New("DeleteBanner: can't get banner id")
	ErrorReadingCreateJson       = errors.New("CreateBanner: can't decode the json")
	ErrorReadingUpdateJson       = errors.New("UpdateBanner: can't decode the json")
)
