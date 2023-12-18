package tables

import (
	"time"

	"github.com/lib/pq"
	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type User struct {
	UUID            string         `gorm:"column:user_pk;primaryKey"`
	FirstName       string         `gorm:"column:first_name"`
	LastName        string         `gorm:"column:last_name"`
	Email           string         `gorm:"column:email"`
	Phone           string         `gorm:"column:phone"`
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
	RoleUUID        string         `gorm:"column:id;primaryKey"`
	RoleDescription *string        `gorm:"column:description"`
	RoleName        string         `gorm:"column:name"`
	RoleScopes      pq.StringArray `gorm:"column:scopes"`
	RoleAuthID      string         `gorm:"column:auth_id"`
}

func (User) TableName() string {
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

func (r *User) ToUserDomain() *domain.User {
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

func (r *User) ToRoleDomain() *domain.Role {
	return &domain.Role{
		UUID:        r.RoleUUID,
		Name:        r.RoleName,
		Description: r.RoleDescription,
		Scopes:      toEnumScopes(r.RoleScopes),
	}
}
