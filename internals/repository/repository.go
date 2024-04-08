package repository

import (
	"encoding/json"
	"log"
	"tools/internals/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	// _"github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddBanner(banner *models.Banner) (int64, int64, error) {
	var id_b, id_tf int64
	bannerContentJson, err := json.Marshal(banner.Content)
	if err != nil {
		log.Printf("Couldn't marshal banner's content: %v", err)
		return 0, 0, nil
	}
	query := insertBanner
	err = r.db.QueryRow(query, bannerContentJson, banner.IsActive).Scan(&id_b)
	if err != nil {
		log.Printf("Couldn't create a new banner: %v", err)
		return 0, 0, err
	}

	query = insertFeatureAndTag
	err1 := r.db.QueryRow(query, id_b, banner.FeatureID, pq.Array(banner.TagIDs)).Scan(&id_tf)
	if err1 != nil {
		log.Printf("Couldn't add new tags and feature: %v", err1)
		return 0, 0, err
	}
	return id_b, id_tf, nil
}
