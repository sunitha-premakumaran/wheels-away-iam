package builders

import (
	"fmt"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type UsersWithRolesBuilder struct {
	size         int
	page         int
	searckKey    *domain.UserSearhKey
	searchString *string
}

func NewUsersWithRolesBuilder(size, page int, searchKey *domain.UserSearhKey, searchString *string) *UsersWithRolesBuilder {
	return &UsersWithRolesBuilder{
		size:         size,
		page:         page,
		searckKey:    searchKey,
		searchString: searchString,
	}
}

func (b *UsersWithRolesBuilder) Build() (string, interface{}) {
	return b.getUsersWithRoles(b.size, b.page, b.searckKey, b.searchString)
}

func (b *UsersWithRolesBuilder) getUsersWithRoles(size, page int, searchKey *domain.UserSearhKey, searchString *string) (string, interface{}) {
	offset := (page - 1) * size
	limit := size
	var params []string

	w := "where deleted_at IS NULL"

	if searchKey != nil && searchString != nil {
		switch *searchKey {
		case domain.Name:
			w = fmt.Sprintf("%s and (UPPER(first_name) LIKE '%?%' or UPPER(last_name) LIKE '%?%')", w)
			params = append(params, *searchString, *searchString)
		case domain.Email:
			w = fmt.Sprintf("%s and UPPER(email) LIKE '%?%'", w)
			params = append(params, *searchString)
		}
	}

	p := fmt.Sprintf("OFFSET %v LIMIT %v", offset, limit)

	s := fmt.Sprintf(`SELECT u.*, r.name as role_name, r.description as role_description, r.scopes as role_scopes, r.role_pk 
	FROM (SELECT user_pk, first_name, last_name, email, primary_phone, alternate_phones, profile_img,
	description, status, auth_id, metadata, created_by, created_at, last_updated_at, last_updated_by,
	deleted_at, deleted_by FROM users %s %s) u 
	left join role_user_mapping m on m.user_id = u.user_pk 
	left join roles r on r.role_pk = m.role_id`, w, p)

	return s, params
}
