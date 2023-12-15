CREATE TABLE users(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(250) NOT NULL,
    primary_phone VARCHAR(250) NOT NULL,
    alternate_phones TEXT[],
    profile_img VARCHAR(250),
    description VARCHAR(250),
    status VARCHAR(250),
    auth_id VARCHAR(250),
    metadata jsonb,
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    last_updated_at timestamp,
    last_updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid
);
