package db

import (
	"context"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"

	"gorm.io/gorm"
)

type (
	User interface {
		// !TODO mockgen doesn't support embedded interface yet
		// !TODO but already discussed in this thread https://github.com/golang/mock/issues/621, lets wait for the release
		// Base[model.UserModel]

		// Base
		Find(ctx context.Context, filterParam abstraction.Filter, search *abstraction.Search) ([]model.UserModel, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id uint64) (*model.UserModel, error)
		FindByCode(ctx context.Context, code string) (*model.UserModel, error)
		FindByName(ctx context.Context, name string) (*model.UserModel, error)
		Create(ctx context.Context, data model.UserModel) (model.UserModel, error)
		Creates(ctx context.Context, data []model.UserModel) ([]model.UserModel, error)
		UpdateByID(ctx context.Context, id uint64, data model.UserModel) (model.UserModel, error)
		DeleteByID(ctx context.Context, id uint64) error
		Count(ctx context.Context) (int64, error)

		// Custom
		FindByEmail(ctx context.Context, email string) (*model.UserModel, error)
	}

	user struct {
		Base[model.UserModel]
	}
)

func NewUser(conn *gorm.DB) User {
	model := model.UserModel{}
	base := NewBase(conn, model, model.TableName())
	return &user{
		base,
	}
}

func (m *user) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	query := m.GetConn(ctx).Debug().Model(model.UserModel{})
	result := new(model.UserModel)
	err := query.Where("email", email).First(result).Error
	if err != nil {
		return nil, m.MaskError(err)
	}
	return result, nil
}
