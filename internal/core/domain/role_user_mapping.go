package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RoleUserMapping struct {
	UUID        string     `validate:"required,uuid4"`
	RoleID      string     `validate:"required,max=150"`
	UserID      string     `validate:"required,max=150"`
	AuthGrantID string     `validate:"required"`
	CreatedAt   time.Time  `validate:"required"`
	CreatedBy   string     `validate:"required,ascii,max=100"`
	DeletedAt   *time.Time `validate:""`
	DeletedBy   *string    `validate:"omitempty,ascii,max=100"`
}

func NewRoleUserMapping(roleID, userID, grantID, createdBy string) (*RoleUserMapping, error) {
	m := &RoleUserMapping{
		RoleID:      roleID,
		UserID:      userID,
		UUID:        uuid.NewString(),
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
		AuthGrantID: grantID,
	}
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return m, nil
}
