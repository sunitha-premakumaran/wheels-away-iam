package tables

import (
	"time"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
)

type RoleUserMapping struct {
	UUID        string     `gorm:"column:role_user_mapping_pk;primaryKey"`
	RoleID      string     `gorm:"column:role_id"`
	UserID      string     `gorm:"column:user_id"`
	AuthGrantID string     `gorm:"column:auth_grant_id"`
	CreatedBy   string     `gorm:"column:created_by"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
	DeletedBy   *string    `gorm:"column:deleted_by"`
}

func (RoleUserMapping) TableName() string {
	return "role_user_mapping"
}

func (m RoleUserMapping) ToDomain() *domain.RoleUserMapping {
	return &domain.RoleUserMapping{
		UUID:        m.UUID,
		RoleID:      m.RoleID,
		UserID:      m.UserID,
		CreatedAt:   m.CreatedAt,
		AuthGrantID: m.AuthGrantID,
		CreatedBy:   m.CreatedBy,
		DeletedAt:   m.DeletedAt,
		DeletedBy:   m.DeletedBy,
	}
}
