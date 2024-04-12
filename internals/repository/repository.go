package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func (r *Repository) FindBanner(input *models.BannerGetRequest) (*models.Banner, error) {
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

func (r *Repository) CreateBanner(input *models.Banner) (*models.BannerID, error) {
	var id_b, id_tf int64
	fmt.Println(input.TagIDs)                             //удалить потом
	bannerContentJson, err := json.Marshal(input.Content) // нужна ли эта ошибка? вынести их по возможности
	if err != nil {
		log.Printf("Couldn't marshal banner's content: %v", err)
		return &models.BannerID{}, err
	}
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Couldn't start a new transaction: %v", err)
		return &models.BannerID{}, err
	}
	query := insertBanner
	err = tx.QueryRow(query, bannerContentJson, input.IsActive).Scan(&id_b)
	if err != nil {
		log.Printf("Couldn't create a new banner: %v", err)
		tx.Rollback()
		return &models.BannerID{}, err
	}

	query = insertFeatureAndTag
	err1 := tx.QueryRow(query, id_b, input.FeatureID, pq.Array(input.TagIDs)).Scan(&id_tf)
	if err1 != nil {
		log.Printf("Couldn't add new tags and feature: %v", err1) //стоит ли обрабатывать перескок id
		tx.Rollback()
		return &models.BannerID{}, err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("Couldn't commit the transaction: %v", err)
		return &models.BannerID{}, err
	}
	return &models.BannerID{BannerId: uint64(id_b)}, nil
}

func (r *Repository) DeleteBanner(input *models.BannerID) error {
	query := deleteBanner
	tx, err := r.db.Begin()
	if err != nil {
		log.Printf("Couldn't start a new transaction: %v", err)
		return err
	}
	_, err = tx.Exec(query, input.BannerId) //тут другой уже(
	if err != nil {
		log.Printf("DeleteBanner: an error occured when deleting: %v", err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("Couldn't commit the transaction: %v", err)
		return err
	}
	return nil
}

func (r *Repository) UpdateBanner(input *models.BannerUpdateRequest) error {
	var deletedFeature uint64
	alreadyDone := false
	tx, err := r.db.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Printf("Couldn't start a new transaction: %v", err)
		return err
	}
	if input.TagIDs != nil {
		query := deleteFeatureTagComb
		err := tx.QueryRow(query, *input.BannerId).Scan(&deletedFeature)
		if err != nil {
			log.Printf("Couldn't delete row when updating: %v", err)
			return err
		}
		if input.FeatureId != nil {
			alreadyDone = true
			query = insertFeatureAndTag
			_, err := r.db.Exec(query, *input.BannerId, *input.FeatureId, pq.Array(input.TagIDs))
			if err != nil {
				log.Printf("Couldn't update row: %v", err)
				return err
			}
		}
	}
	if input.FeatureId != nil && !alreadyDone {
		query := updateFeature
		_, err := r.db.Exec(query, *input.FeatureId, *input.BannerId)
		if err != nil {
			log.Printf("Couldn't delete row when updating: %v", err)
			return err
		}
	}
	if input.Content != nil {
		query := updateContent
		_, err := r.db.Exec(query, *input.Content, *input.BannerId)
		if err != nil {
			log.Printf("Couldn't update the content: %v", err)
			return err
		}
	}
	if input.IsActive != nil {
		query := updateIsActive
		_, err := r.db.Exec(query, *input.IsActive, *input.BannerId)
		if err != nil {
			log.Printf("Couldn't update the isActive status: %v", err)
			return err
		}
	}
	tx.Commit()
	return nil
}

func (r *Repository) GetFilteredBanner(input *models.BannerGetAdminRequest) ([]models.Banner, error) {

	var count int
	query := getBannerAdmin
	fmt.Println(input.FeatureID)
	fmt.Println(input.TagID)
	fmt.Println(input.Offset)
	fmt.Println(input.Limit)
	rows, err := r.db.Queryx(query, input.FeatureID, input.TagID, input.Limit, input.Offset)
	if err != nil {
		log.Printf("GetFilteredBanner: couldn't execute queryx: %v", err)
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	banners := make([]models.Banner, 0)
	for rows.Next() {
		var banner models.Banner
		if err := rows.Scan(&banner.ID, pq.Array(&banner.TagIDs), &banner.Content, &banner.FeatureID, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt); err != nil {
			log.Printf("GetFilteredBanner: error scaning a row: %v", err)
			return banners, nil
		}
		banners = append(banners, banner)
		count += 1
	}
	fmt.Println(count)

	if err = rows.Err(); err != nil {
		log.Printf("GetFilteredBanner: couldn't execute the filter query: %v", err)
		return nil, err
	}

	return banners, nil
}
