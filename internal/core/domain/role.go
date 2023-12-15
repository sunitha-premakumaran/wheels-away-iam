package domain

import (
	"time"

	"github.com/sunitha/wheels-away-iam/internal/core/enums"
)

type Role struct {
	UUID          string
	Name          string
	Description   string
	Scopes        []enums.UserScope
	AuthID        string
	CreatedBy     string
	CreatedAt     time.Time
	LastUpdatedAt *time.Time
	LastUpdatedBy *string
	DeletedAt     *time.Time
	DeletedBy     *string
}
