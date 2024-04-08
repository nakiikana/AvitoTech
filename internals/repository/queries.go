package repository

const (
	insertBanner = `
		INSERT INTO banner (content, is_active)
		VALUES ($1, $2)
		RETURNING id
	`
	insertFeatureAndTag = `
		INSERT INTO tag_feature_banner (banner_id, feature_id, tag_id)
		SELECT $1, $2, tag
		FROM unnest($3::bigint[]) as tag
		RETURNING id
	`
	deleteBanner = `
		DELETE FROM banner WHERE id = $1 RETURNING id
	`
	deleteFeatureTagComb = `
		DELETE FROM tag_feature_banner WHERE tag_id = $1 AND feature_id = $2
	`
	getBanner = `
		SELECT content FROM banner b
		INNER JOIN (SELECT banner_id FROM tag_feature_banner 
		WHERE feature_id = $1 AND tag_id = $2) AS t ON t.banner_id = b.banner_id
		WHERE is_active = true
	`
)
