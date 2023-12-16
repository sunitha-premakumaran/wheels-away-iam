package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type Role struct {
	UUID          string            `validate:"required,uuid4"`
	Name          string            `validate:"required,max=150"`
	Description   *string           `validate:"omitempty,ascii,max=250"`
	AuthID        string            `validate:"required,ascii,max=250"`
	Scopes        []enums.UserScope `validate:"required"`
	CreatedAt     time.Time         `validate:"required"`
	CreatedBy     string            `validate:"required,ascii,max=100"`
	LastUpdatedAt *time.Time        `validate:""`
	LastUpdatedBy *string           `validate:"omitempty,ascii,max=100"`
	DeletedAt     *time.Time        `validate:""`
	DeletedBy     *string           `validate:"omitempty,ascii,max=100"`
}

func NewRole(name string, description *string, scopes []enums.UserScope, authID, createdBy string) (*Role, error) {
	r := &Role{
		UUID:        uuid.NewString(),
		Name:        name,
		Description: description,
		Scopes:      scopes,
		AuthID:      authID,
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
	}
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return r, nil
}
