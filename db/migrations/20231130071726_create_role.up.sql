CREATE TABLE roles(
    id  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(250) NOT NULL,
	description VARCHAR(250),
	associated_scopes  TEXT[],
	auth_id  VARCHAR(250),
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    last_updated_at timestamp,
    last_updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid
)