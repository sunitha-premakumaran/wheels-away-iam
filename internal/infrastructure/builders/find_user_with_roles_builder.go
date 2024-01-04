package builders

import (
	"fmt"
)

type FindUserWithRolesBuilder struct {
	userID string
}

func NewFindUserWithRolesBuilder(userID string) *FindUserWithRolesBuilder {
	return &FindUserWithRolesBuilder{
		userID: userID,
	}
}

func (b *FindUserWithRolesBuilder) Build() string {
	return b.getFindUserWithRoles(b.userID)
}

func (b *FindUserWithRolesBuilder) getFindUserWithRoles(userId string) string {

	w := fmt.Sprintf("where deleted_at IS NULL and user_pk = '%s'", userId)

	s := fmt.Sprintf(`SELECT u.*, r.name as role_name, r.description as role_description, r.scopes as role_scopes, r.role_pk 
	FROM (SELECT user_pk, first_name, last_name, email, primary_phone, alternate_phones, profile_img,
	description, status, auth_id, metadata, created_by, created_at, lastupdated_at, lastupdated_by,
	deleted_at, deleted_by FROM users %s) u 
	left join role_user_mapping m on m.user_id = u.user_pk 
	left join roles r on r.role_pk = m.role_id`, w)

	return s
}
