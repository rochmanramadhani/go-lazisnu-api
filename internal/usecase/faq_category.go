package usecase

import (
	"context"
	"github.com/google/uuid"
	"strings"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/trxmanager"
)

type FaqCategory interface {
	Find(ctx context.Context, filterParam abstraction.Filter) ([]dto.FaqCategoryResponse, abstraction.PaginationInfo, error)
	FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.FaqCategoryResponse, error)
	Create(ctx context.Context, payload dto.CreateFaqCategoryRequest) (dto.FaqCategoryResponse, error)
	Update(ctx context.Context, payload dto.UpdateFaqCategoryRequest) (dto.FaqCategoryResponse, error)
	Delete(ctx context.Context, payload dto.ByIDRequest) (dto.FaqCategoryResponse, error)
}

type faqCategory struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewFaqCategory(cfg *config.Configuration, f repository.Factory) FaqCategory {
	return &faqCategory{cfg, f}
}

func (u *faqCategory) Find(ctx context.Context, filterParam abstraction.Filter) (result []dto.FaqCategoryResponse, pagination abstraction.PaginationInfo, err error) {
	var search *abstraction.Search
	if filterParam.Search != "" {
		searchQuery := "lower(code) LIKE ? OR lower(name) LIKE ?"
		searchVal := "%" + strings.ToLower(filterParam.Search) + "%"
		search = &abstraction.Search{
			Query: searchQuery,
			Args:  []interface{}{searchVal, searchVal},
		}
	}

	faqCategories, info, err := u.Repo.FaqCategory.Find(ctx, filterParam, search)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	pagination = *info

	faqs, _, err := u.Repo.Faq.Find(ctx, abstraction.Filter{}, nil)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	for _, faqCategory := range faqCategories {
		result = append(result, dto.FaqCategoryResponse{
			ID:     faqCategory.ID,
			UUID:   faqCategory.UUID,
			Name:   faqCategory.Name,
			Status: faqCategory.Status,
			Faqs:   nil,
		})

		for _, faq := range faqs {
			if faq.FaqCategoryID == faqCategory.ID {
				result[len(result)-1].Faqs = append(result[len(result)-1].Faqs, dto.FaqResponse{
					ID:       faq.ID,
					UUID:     faq.UUID,
					Question: faq.Question,
					Answer:   faq.Answer,
				})
			}
		}
	}

	return result, pagination, nil
}

func (u *faqCategory) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.FaqCategoryResponse, error) {
	var result dto.FaqCategoryResponse

	faqCategory, err := u.Repo.FaqCategory.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	faq, _, err := u.Repo.Faq.Find(ctx, abstraction.Filter{}, nil)
	if err != nil {
		return result, err
	}

	result = dto.FaqCategoryResponse{
		ID:     faqCategory.ID,
		UUID:   faqCategory.UUID,
		Name:   faqCategory.Name,
		Status: faqCategory.Status,
		Faqs:   nil,
	}

	for _, faq := range faq {
		if faq.FaqCategoryID == faqCategory.ID {
			result.Faqs = append(result.Faqs, dto.FaqResponse{
				ID:       faq.ID,
				UUID:     faq.UUID,
				Question: faq.Question,
				Answer:   faq.Answer,
			})
		}
	}

	return result, nil
}

func (u *faqCategory) Create(ctx context.Context, payload dto.CreateFaqCategoryRequest) (result dto.FaqCategoryResponse, err error) {
	var (
		faqCategory = model.FaqCategoryModel{
			FaqCategoryEntity: model.FaqCategoryEntity{
				UUID: uuid.New(),
				Name: payload.Name,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		faqCategory, err = u.Repo.FaqCategory.Create(ctx, faqCategory)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqCategoryResponse{
		ID:     faqCategory.ID,
		UUID:   faqCategory.UUID,
		Name:   faqCategory.Name,
		Status: faqCategory.Status,
		Faqs:   nil,
	}

	return result, nil
}

func (u *faqCategory) Update(ctx context.Context, payload dto.UpdateFaqCategoryRequest) (result dto.FaqCategoryResponse, err error) {
	var (
		faqCategory = &model.FaqCategoryModel{
			FaqCategoryEntity: model.FaqCategoryEntity{
				Name:   payload.Name,
				Status: payload.Status,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		faqCategory, err = u.Repo.FaqCategory.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.Name != "" {
			faqCategory.Name = payload.Name
		}

		if payload.Status != 0 {
			faqCategory.Status = payload.Status
		}

		_, err = u.Repo.FaqCategory.UpdateByID(ctx, payload.ID, *faqCategory)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqCategoryResponse{
		ID:     faqCategory.ID,
		UUID:   faqCategory.UUID,
		Name:   faqCategory.Name,
		Status: faqCategory.Status,
		Faqs:   nil,
	}

	return result, nil
}

func (u *faqCategory) Delete(ctx context.Context, payload dto.ByIDRequest) (result dto.FaqCategoryResponse, err error) {
	var data *model.FaqCategoryModel

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		data, err = u.Repo.FaqCategory.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.FaqCategory.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqCategoryResponse{
		ID:     data.ID,
		UUID:   data.UUID,
		Name:   data.Name,
		Status: data.Status,
		Faqs:   nil,
	}

	return result, nil
}
