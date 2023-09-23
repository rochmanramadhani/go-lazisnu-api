package db

import (
	"context"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"

	"gorm.io/gorm"
)

type (
	Faq interface {
		// !TODO mockgen doesn't support embedded interface yet
		// !TODO but already discussed in this thread https://github.com/golang/mock/issues/621, lets wait for the release
		// Base[model.FaqModel]

		// Base
		Find(ctx context.Context, filterParam abstraction.Filter, search *abstraction.Search) ([]model.FaqModel, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id uint64) (*model.FaqModel, error)
		FindByCode(ctx context.Context, code string) (*model.FaqModel, error)
		FindByName(ctx context.Context, name string) (*model.FaqModel, error)
		Create(ctx context.Context, data model.FaqModel) (model.FaqModel, error)
		Creates(ctx context.Context, data []model.FaqModel) ([]model.FaqModel, error)
		UpdateByID(ctx context.Context, id uint64, data model.FaqModel) (model.FaqModel, error)
		DeleteByID(ctx context.Context, id uint64) error
		Count(ctx context.Context) (int64, error)
	}

	faq struct {
		Base[model.FaqModel]
	}
)

func NewFaq(conn *gorm.DB) Faq {
	model := model.FaqModel{}
	base := NewBase(conn, model, model.TableName())
	return &faq{
		base,
	}
}
