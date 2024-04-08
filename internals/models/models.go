package models

import "time"

type Banner struct {
	ID        uint64
	Content   string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	TagIDs    []uint64
	FeatureID uint64
}
