package domain

type DecoratedUser struct {
	User      *User
	UserRoles []*Role
}
