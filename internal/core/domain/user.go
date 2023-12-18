package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type User struct {
	UUID          string           `validate:"required,uuid4"`
	FirstName     string           `validate:"required,max=150"`
	LastName      string           `validate:"required,max=150"`
	Password      string           `validate:"required,max=150"`
	Email         string           `validate:"required,max=250"`
	Phone         string           `validate:"required,max=10"`
	ProfileImage  *string          `validate:"omitempty,ascii,max=250"`
	Description   *string          `validate:"omitempty,ascii,max=250"`
	Status        enums.UserStatus `validate:"required"`
	AuthID        string           `validate:"required,ascii,max=250"`
	Metadata      interface{}      `validate:"omitempty"`
	CreatedAt     time.Time        `validate:"required"`
	CreatedBy     string           `validate:"required,ascii,max=100"`
	LastUpdatedAt *time.Time       `validate:""`
	LastUpdatedBy *string          `validate:"omitempty,ascii,max=100"`
	DeletedAt     *time.Time       `validate:""`
	DeletedBy     *string          `validate:"omitempty,ascii,max=100"`
}

func NewUser(firstName, lastName, email, phone string, profileImageUrl, description *string,
	status enums.UserStatus, authID, createdBy string) (*User, error) {
	u := &User{
		UUID:         uuid.NewString(),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Phone:        phone,
		ProfileImage: profileImageUrl,
		Description:  description,
		Status:       status,
		AuthID:       authID,
		CreatedAt:    time.Now(),
		CreatedBy:    createdBy,
	}
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return u, nil
}
