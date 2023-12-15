CREATE TABLE role_user_map(
    role_id_user_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id uuid REFERENCES roles (id),
    user_id uuid REFERENCES users (id)
)