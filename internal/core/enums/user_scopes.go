package enums

type UserScope string

const (
	USERS_READ  UserScope = "users.read"
	USERS_WRITE UserScope = "users.write"
	ROLES_READ  UserScope = "roles.read"
	ROLES_WRITE UserScope = "roles.write"
)
