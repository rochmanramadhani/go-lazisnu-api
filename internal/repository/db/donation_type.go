package db

import (
	"context"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"

	"gorm.io/gorm"
)

type (
	DonationType interface {
		// !TODO mockgen doesn't support embedded interface yet
		// !TODO but already discussed in this thread https://github.com/golang/mock/issues/621, lets wait for the release
		// Base[model.DonationTypeModel]

		// Base
		Find(ctx context.Context, filterParam abstraction.Filter, search *abstraction.Search) ([]model.DonationTypeModel, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id uint64) (*model.DonationTypeModel, error)
		FindByCode(ctx context.Context, code string) (*model.DonationTypeModel, error)
		FindByName(ctx context.Context, name string) (*model.DonationTypeModel, error)
		Create(ctx context.Context, data model.DonationTypeModel) (model.DonationTypeModel, error)
		Creates(ctx context.Context, data []model.DonationTypeModel) ([]model.DonationTypeModel, error)
		UpdateByID(ctx context.Context, id uint64, data model.DonationTypeModel) (model.DonationTypeModel, error)
		DeleteByID(ctx context.Context, id uint64) error
		Count(ctx context.Context) (int64, error)
	}

	donationType struct {
		Base[model.DonationTypeModel]
	}
)

func NewDonationType(conn *gorm.DB) DonationType {
	model := model.DonationTypeModel{}
	base := NewBase(conn, model, model.TableName())
	return &donationType{
		base,
	}
}
