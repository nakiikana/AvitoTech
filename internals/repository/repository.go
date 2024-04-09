package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"tools/internals/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Couldn't start a new transaction: %v", err)
		return 0, 0, err
	}
	query := insertBanner
	err = tx.QueryRow(query, bannerContentJson, banner.IsActive).Scan(&id_b)
	if err != nil {
		log.Printf("Couldn't create a new banner: %v", err)
		tx.Rollback()
		return 0, 0, err
	}

	query = insertFeatureAndTag
	err1 := tx.QueryRow(query, id_b, banner.FeatureID, pq.Array(banner.TagIDs)).Scan(&id_tf)
	if err1 != nil {
		log.Printf("Couldn't add new tags and feature: %v", err1)
		tx.Rollback()
		return 0, 0, err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("Couldn't commit the transaction: %v", err)
		return 0, 0, err
	}
	return id_b, id_tf, nil
}

func (r *Repository) FindBanner(input *models.BannerGetMethod) (*models.Banner, error) {
	content := &models.Banner{}
	query := getBanner
	ans := r.db.QueryRow(query, input.FeatureID, input.TagID)
	if err := ans.Scan(&content.Content); err != nil {
		if err = sql.ErrNoRows; err != nil {

			log.Printf("FindBanner: no rows found: %v tag_id: %d feature_id: %d", err, input.TagID, input.FeatureID)
			return &models.Banner{}, err
		} else {
			log.Printf("FindBanner: an error occured when looking for a banner: %v", err)
			return &models.Banner{}, err
		}
	}
	return content, nil

}
