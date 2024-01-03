package queries

import (
	"time"

	"github.com/lib/pq"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type UserWithRolesRow struct {
	UUID            string         `gorm:"column:user_pk;primaryKey"`
	FirstName       string         `gorm:"column:first_name"`
	LastName        string         `gorm:"column:last_name"`
	Email           string         `gorm:"column:email"`
	Phone           string         `gorm:"column:primary_phone"`
	ProfileImage    *string        `gorm:"column:profile_img"`
	Description     *string        `gorm:"column:description"`
	Status          string         `gorm:"column:status"`
	AuthID          string         `gorm:"column:auth_id"`
	CreatedBy       string         `gorm:"column:created_by"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	LastUpdatedAt   *time.Time     `gorm:"column:lastupdated_at"`
	LastUpdatedBy   *string        `gorm:"column:lastupdated_by"`
	DeletedAt       *time.Time     `gorm:"column:deleted_at"`
	DeletedBy       *string        `gorm:"column:deleted_by"`
	RoleUUID        string         `gorm:"column:role_pk;primaryKey"`
	RoleDescription *string        `gorm:"column:role_description"`
	RoleName        string         `gorm:"column:role_name"`
	RoleScopes      pq.StringArray `gorm:"column:role_scopes;type:text[]"`
}

func (UserWithRolesRow) TableName() string {
	return "users"
}

func toEnumStatus(status string) enums.UserStatus {
	switch status {
	case "ACTIVE":
		return enums.ACTIVE
	case "IN_ACTIVE":
		return enums.INACTIVE
	}
	return ""
}

func (r *UserWithRolesRow) ToUserDomain() *domain.User {
	return &domain.User{
		UUID:          r.UUID,
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		Email:         r.Email,
		Phone:         r.Phone,
		ProfileImage:  r.ProfileImage,
		Description:   r.Description,
		Status:        toEnumStatus(r.Status),
		AuthID:        r.AuthID,
		CreatedBy:     r.CreatedBy,
		CreatedAt:     r.CreatedAt,
		LastUpdatedAt: r.LastUpdatedAt,
		LastUpdatedBy: r.LastUpdatedBy,
		DeletedAt:     r.DeletedAt,
		DeletedBy:     r.DeletedBy,
	}
}

func toEnumScopes(scopes []string) []enums.UserScope {
	us := make([]enums.UserScope, 0, len(scopes))
	for _, s := range us {
		switch s {
		case "roles.read":
			us = append(us, enums.ROLES_READ)
		case "roles.write":
			us = append(us, enums.ROLES_WRITE)
		case "users.read":
			us = append(us, enums.USERS_READ)
		case "users.write":
			us = append(us, enums.USERS_WRITE)
		}
	}
	return us
}

func (r *UserWithRolesRow) ToRoleDomain() *domain.Role {
	return &domain.Role{
		UUID:        r.RoleUUID,
		Name:        r.RoleName,
		Description: r.RoleDescription,
		Scopes:      toEnumScopes(r.RoleScopes),
	}
}
