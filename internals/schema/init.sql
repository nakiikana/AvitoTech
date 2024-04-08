CREATE TABLE IF NOT EXISTS banner (
		id bigserial primary key not null, 
		content JSONB not null,
		is_active boolean not null, 
		created_at timestamp DEFAULT NOW(),
		updated_at timestamp DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tag_feature_banner (
		id bigserial primary key not null,
		banner_id bigint not null references banner (id) on delete cascade, 
		tag_id bigint not null, 
		feature_id bigint not null, 
		constraint  unique_data UNIQUE (tag_id, feature_id)
);