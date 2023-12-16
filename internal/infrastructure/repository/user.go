package repository

import (
	"context"
	"errors"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/tables"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		gormDB: db,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*domain.DecoratedUser, error) {
	return r.getUsers()
}

func (r *UserRepository) getUsers() ([]*domain.DecoratedUser, error) {
	var users []*tables.User
	result := r.gormDB.Model(&tables.User{}).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	var mu map[string]*domain.DecoratedUser
	for _, u := range users {
		if us, ok := mu[u.UUID]; ok {
			us.UserRoles = append(us.UserRoles, u.ToRoleDomain())

		} else {
			mu[u.UUID] = &domain.DecoratedUser{
				User:      u.ToUserDomain(),
				UserRoles: []*domain.Role{u.ToRoleDomain()},
			}
		}
	}
	var du []*domain.DecoratedUser
	for _, u := range mu {
		du = append(du, u)

	}
	return du, nil
}

func (r *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	userToUpsert := mapDomainToTable(user)
	return r.saveUser(ctx, userToUpsert)
}

func (r *UserRepository) saveUser(ctx context.Context, user *tables.User) error {

	tx := r.gormDB.WithContext(ctx)

	result := tx.Session(&gorm.Session{FullSaveAssociations: true}).Clauses(clause.OnConflict{
		Where: clause.Where{Exprs: []clause.Expression{
			clause.Lt{
				Column: `"users"."lastupdated_at"`,
				Value:  user.LastUpdatedAt,
			},
			clause.Or(
				clause.Eq{
					Column: `"users"."lastupdated_at"`,
					Value:  nil,
				},
			),
		}},
		UpdateAll: true,
	}).Create(&user)

	return result.Error
}

func mapDomainToTable(user *domain.User) *tables.User {
	return &tables.User{
		UUID:          user.UUID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Phone:         user.Phone,
		ProfileImage:  user.ProfileImage,
		Description:   user.Description,
		Status:        string(user.Status),
		AuthID:        user.AuthID,
		CreatedBy:     user.CreatedBy,
		CreatedAt:     user.CreatedAt,
		LastUpdatedAt: user.LastUpdatedAt,
		LastUpdatedBy: user.LastUpdatedBy,
		DeletedAt:     user.DeletedAt,
		DeletedBy:     user.DeletedBy,
	}
}
