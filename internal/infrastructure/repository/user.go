package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sunitha/wheels-away-iam/internal/core/domain"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/builders"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/models/queries"
	"github.com/sunitha/wheels-away-iam/internal/infrastructure/models/tables"
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

func (r *UserRepository) GetUsers(ctx context.Context, page, size int,
	searchKey *domain.UserSearhKey, searchString *string) ([]*domain.DecoratedUser, *domain.PageInfo, error) {
	users, err := r.getUsers(page, size, searchKey, searchString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting users: %w", err)
	}
	pageInfo, err := r.getUsersCount(page, size, searchKey, searchString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user_count: %w", err)
	}
	return users, pageInfo, nil
}

func (r *UserRepository) getUsers(page, size int, searchKey *domain.UserSearhKey,
	searchString *string) ([]*domain.DecoratedUser, error) {
	var users []*queries.UserWithRolesRow
	builder := builders.NewUsersWithRolesBuilder(page, size, searchKey, searchString)
	rawSQL, _ := builder.Build()
	result := r.gormDB.Raw(rawSQL).Find(&users)
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

func (r *UserRepository) getUsersCount(page, size int, searchKey *domain.UserSearhKey,
	searchString *string) (*domain.PageInfo, error) {
	var count int64
	query := r.gormDB.Model(&tables.User{}).Where("deleted_at IS NULL")
	if searchKey != nil && searchString != nil {
		switch *searchKey {
		case domain.Name:
			query = query.Where("upper(first_name) LIKE upper(%?%)", searchString).Or("upper(last_name) LIKE upper(%?%)", searchString)
		case domain.Email:
			query = query.Where("upper(email) LIKE upper(%?%)", searchString)
		}
	}

	result := query.Count(&count)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return toPageInfo(int(count), page, size), nil
}

func toPageInfo(count, pageNumber, pageSize int) *domain.PageInfo {
	totalItems := count
	totalPages := (count + pageSize - 1) / pageSize

	return &domain.PageInfo{
		CurrentPage: pageNumber,
		PageSize:    pageSize,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}
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
