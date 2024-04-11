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
		DELETE FROM tag_feature_banner WHERE banner_id = $1 RETURNING feature_id
	`
	getBanner = `
		SELECT content FROM banner b
		INNER JOIN (SELECT banner_id FROM tag_feature_banner 
		WHERE feature_id = $1 AND tag_id = $2) AS t ON t.banner_id = b.id
		WHERE is_active = true
	`
	updateFeature = `
		UPDATE tag_feature_banner SET feature_id = $1 WHERE banner_id = $2
	`
	updateIsActive = `
		UPDATE banner SET is_active = $1 WHERE id = $2
	`
	updateContent = `
		UPDATE banner SET content = $1 WHERE id = $2
	`
	getBannerAdmin = `
		SELECT b.id, ARRAY_AGG(ftb.tag_id)::bigint[] AS tag_ids, ftb.feature_id, b.is_active, b.created_at, b.updated_at 
		FROM banner b
		INNER JOIN tag_feature_banner ftb ON ftb.banner_id = b.id
		WHERE b.id IN (
        SELECT DISTINCT banner_id 
        FROM features_tags_banner 
        WHERE 
            ($1::bigint IS NOT NULL AND feature_id = $1) OR 
            ($2::bigint IS NOT NULL AND tag_id = $2))
		GROUP BY b.id, ftb.feature_id
		LIMIT $3 
		OFFSET $4
	`
)
