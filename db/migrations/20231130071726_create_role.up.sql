CREATE TABLE roles(
    role_pk  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR(250) NOT NULL,
	description VARCHAR(250),
	scopes  TEXT[],
    auth_key VARCHAR(150) UNIQUE NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    last_updated_at timestamp,
    last_updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid
)