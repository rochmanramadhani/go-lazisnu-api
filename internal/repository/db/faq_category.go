package db

import (
	"context"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"

	"gorm.io/gorm"
)

type (
	FaqCategory interface {
		// !TODO mockgen doesn't support embedded interface yet
		// !TODO but already discussed in this thread https://github.com/golang/mock/issues/621, lets wait for the release
		// Base[model.FaqCategoryModel]

		// Base
		Find(ctx context.Context, filterParam abstraction.Filter, search *abstraction.Search) ([]model.FaqCategoryModel, *abstraction.PaginationInfo, error)
		FindByID(ctx context.Context, id uint64) (*model.FaqCategoryModel, error)
		FindByCode(ctx context.Context, code string) (*model.FaqCategoryModel, error)
		FindByName(ctx context.Context, name string) (*model.FaqCategoryModel, error)
		Create(ctx context.Context, data model.FaqCategoryModel) (model.FaqCategoryModel, error)
		Creates(ctx context.Context, data []model.FaqCategoryModel) ([]model.FaqCategoryModel, error)
		UpdateByID(ctx context.Context, id uint64, data model.FaqCategoryModel) (model.FaqCategoryModel, error)
		DeleteByID(ctx context.Context, id uint64) error
		Count(ctx context.Context) (int64, error)
	}

	faqCategory struct {
		Base[model.FaqCategoryModel]
	}
)

func NewFaqCategory(conn *gorm.DB) FaqCategory {
	model := model.FaqCategoryModel{}
	base := NewBase(conn, model, model.TableName())
	return &faqCategory{
		base,
	}
}
