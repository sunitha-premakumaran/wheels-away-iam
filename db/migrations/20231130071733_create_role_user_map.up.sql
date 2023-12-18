CREATE TABLE role_user_mapping(
    role_user_mapping_pk uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id uuid REFERENCES roles (role_pk),
    user_id uuid REFERENCES users (user_pk),
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp,
    deleted_by uuid
)