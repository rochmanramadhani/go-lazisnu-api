package db

import (
	"context"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"

	"gorm.io/gorm"
)

type (
	UserProfile interface {
		// !TODO mockgen doesn't support embedded interface yet
		// !TODO but already discussed in this thread https://github.com/golang/mock/issues/621, lets wait for the release
		// Base[model.UserProfileModel]

		// Base
		Find(ctx context.Context, filterParam abstraction.Filter, search *abstraction.Search) ([]model.UserProfileModel, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id uint64) (*model.UserProfileModel, error)
		FindByCode(ctx context.Context, code string) (*model.UserProfileModel, error)
		FindByName(ctx context.Context, name string) (*model.UserProfileModel, error)
		Create(ctx context.Context, data model.UserProfileModel) (model.UserProfileModel, error)
		Creates(ctx context.Context, data []model.UserProfileModel) ([]model.UserProfileModel, error)
		UpdateByID(ctx context.Context, id uint64, data model.UserProfileModel) (model.UserProfileModel, error)
		DeleteByID(ctx context.Context, id uint64) error
		Count(ctx context.Context) (int64, error)

		// Custom
		FindByUserID(ctx context.Context, userID uint64) (*model.UserProfileModel, error)
		UpdateByUserID(ctx context.Context, userID uint64, data model.UserProfileModel) (model.UserProfileModel, error)
		DeleteByUserID(ctx context.Context, userID uint64) error
	}

	userProfile struct {
		Base[model.UserProfileModel]
	}
)

func NewUserProfile(conn *gorm.DB) UserProfile {
	model := model.UserProfileModel{}
	base := NewBase(conn, model, model.TableName())
	return &userProfile{
		base,
	}
}

func (m *userProfile) FindByUserID(ctx context.Context, userID uint64) (*model.UserProfileModel, error) {
	query := m.GetConn(ctx).Debug().Model(model.UserProfileModel{})
	result := new(model.UserProfileModel)
	err := query.Where("user_id", userID).First(result).Error
	if err != nil {
		return nil, m.MaskError(err)
	}
	return result, nil
}

func (m *userProfile) UpdateByUserID(ctx context.Context, userID uint64, data model.UserProfileModel) (model.UserProfileModel, error) {
	err := m.GetConn(ctx).Debug().Model(&data).Where("user_id = ?", userID).Updates(&data).Error
	return data, err
}

func (m *userProfile) DeleteByUserID(ctx context.Context, userID uint64) error {
	return m.GetConn(ctx).Where("user_id = ?", userID).Delete(&model.UserProfileModel{}).Error
}
