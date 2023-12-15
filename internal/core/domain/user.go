package domain

import (
	"time"

	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type UserMetadata interface {
}

type User struct {
	UUID          string
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	ProfileImage  string
	Description   string
	Status        enums.UserStatus
	AuthID        string
	Metadata      UserMetadata
	CreatedBy     string
	CreatedAt     time.Time
	LastUpdatedAt *time.Time
	LastUpdatedBy *string
	DeletedAt     *time.Time
	DeletedBy     *string
}
