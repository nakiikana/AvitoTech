package models

import "time"

type Banner struct {
	ID        uint64
	Content   string
	IsActive  bool
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
