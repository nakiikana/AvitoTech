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
	SELECT banner.id, array_agg(ftb.tag_id)::bigint[] as tag_ids, banner.content, ftb.feature_id, banner.is_active,
	banner.created_at, banner.updated_at FROM banner
	INNER JOIN tag_feature_banner as ftb ON (ftb.banner_id = banner.id)	
	WHERE banner.id in (
	(
	 SELECT DISTINCT banner_id FROM tag_feature_banner
	 WHERE (CASE WHEN $1::bigint IS NOT NULL THEN feature_id = $1 ELSE true END)
	   and (CASE WHEN $2::bigint IS NOT NULL THEN tag_id = $2 ELSE true END) 
 	)
	)
	GROUP BY banner.id, ftb.feature_id
	LIMIT $3 OFFSET $4
	`
)
