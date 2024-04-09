package models

import "time"

type Banner struct {
	ID        uint64
	Content   string `json: "context"`
	IsActive  bool   `json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
	TagIDs    []uint64 `json:"tag_id"`
	FeatureID uint64   `json:"feature_id"`
}

type BannerGetMethod struct {
	TagID           uint64 // почему не массив
	FeatureID       uint64
	UseLastRevision bool
	// !Token
}

type InsertedBannerResponse struct {
	BannerId uint64 `json:"banner_id"`
}
