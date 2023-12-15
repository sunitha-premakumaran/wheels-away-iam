package tables

import (
	"time"

	"github.com/lib/pq"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type Role struct {
	UUID          string         `gorm:"column:id;primaryKey"`
	Description   string         `gorm:"column:description"`
	Name          string         `gorm:"column:name"`
	Scopes        pq.StringArray `gorm:"column:scopes"`
	AuthID        string         `gorm:"column:auth_id"`
	CreatedBy     string         `gorm:"column:created_by"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	LastUpdatedAt *time.Time     `gorm:"column:lastupdated_at"`
	LastUpdatedBy *string        `gorm:"column:lastupdated_by"`
	DeletedAt     *time.Time     `gorm:"column:deleted_at"`
	DeletedBy     *string        `gorm:"column:deleted_by"`
}

func (Role) TableName() string {
	return "roles"
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

func (r *Role) ToRoleDomain() *domain.Role {
	return &domain.Role{
		UUID:          r.UUID,
		Name:          r.Name,
		Description:   r.Description,
		Scopes:        toEnumScopes(r.Scopes),
		AuthID:        r.AuthID,
		CreatedBy:     r.CreatedBy,
		CreatedAt:     r.CreatedAt,
		LastUpdatedAt: r.LastUpdatedAt,
		LastUpdatedBy: r.LastUpdatedBy,
		DeletedAt:     r.DeletedAt,
		DeletedBy:     r.DeletedBy,
	}
}
