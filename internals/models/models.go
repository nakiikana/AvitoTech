package models

import (
	"encoding/json"
	"time"
)

type Banner struct {
	ID        uint64          `json: "banner_id"`
	Content   json.RawMessage `json: "content"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
	TagIDs    []uint64 `json:"tag_ids"`
	FeatureID uint64   `json:"feature_id"`
}

type BannerUpdateRequest struct {
	BannerId  *uint64
	Content   *json.RawMessage `json: "content"`
	TagIDs    *[]int64         `json:"tag_ids"`
	FeatureId *uint64          `json:"feature_id"`
	IsActive  *bool            `json:"is_active"`
}

type BannerGetRequest struct {
	TagID           uint64
	FeatureID       uint64
	UseLastRevision bool
	// !Token
}

type BannerGetAdminRequest struct {
	TagID     uint64
	FeatureID uint64
	Limit     int
	Offset    int
}

type BannerID struct {
	BannerId uint64 `json:"banner_id"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}
