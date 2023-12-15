CREATE TABLE role_user_map(
    role_id_user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id uuid REFERENCES roles (role_id),
    user_id uuid REFERENCES users (user_id),
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    deleted_at timestamp,
    deleted_by uuid
)