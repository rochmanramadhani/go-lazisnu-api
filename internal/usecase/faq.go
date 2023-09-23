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

type Faq interface {
	Find(ctx context.Context, filterParam abstraction.Filter) ([]dto.FaqResponse, abstraction.PaginationInfo, error)
	FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.FaqResponse, error)
	Create(ctx context.Context, payload dto.CreateFaqRequest) (dto.FaqResponse, error)
	Update(ctx context.Context, payload dto.UpdateFaqRequest) (dto.FaqResponse, error)
	Delete(ctx context.Context, payload dto.ByIDRequest) (dto.FaqResponse, error)
}

type faq struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewFaq(cfg *config.Configuration, f repository.Factory) Faq {
	return &faq{cfg, f}
}

func (u *faq) Find(ctx context.Context, filterParam abstraction.Filter) (result []dto.FaqResponse, pagination abstraction.PaginationInfo, err error) {
	var search *abstraction.Search
	if filterParam.Search != "" {
		searchQuery := "lower(code) LIKE ? OR lower(name) LIKE ?"
		searchVal := "%" + strings.ToLower(filterParam.Search) + "%"
		search = &abstraction.Search{
			Query: searchQuery,
			Args:  []interface{}{searchVal, searchVal},
		}
	}

	faqs, info, err := u.Repo.Faq.Find(ctx, filterParam, search)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	pagination = *info

	for _, faq := range faqs {
		result = append(result, dto.FaqResponse{
			ID:       faq.ID,
			UUID:     faq.UUID,
			Question: faq.Question,
			Answer:   faq.Answer,
			Status:   faq.Status,
		})
	}

	return result, pagination, nil
}

func (u *faq) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.FaqResponse, error) {
	var result dto.FaqResponse

	faq, err := u.Repo.Faq.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	result = dto.FaqResponse{
		ID:       faq.ID,
		UUID:     faq.UUID,
		Question: faq.Question,
		Answer:   faq.Answer,
		Status:   faq.Status,
	}

	return result, nil
}

func (u *faq) Create(ctx context.Context, payload dto.CreateFaqRequest) (result dto.FaqResponse, err error) {
	var (
		faq = model.FaqModel{
			FaqEntity: model.FaqEntity{
				UUID:          uuid.New(),
				Question:      payload.Question,
				Answer:        payload.Answer,
				FaqCategoryID: payload.FaqCategoryID,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		faq, err = u.Repo.Faq.Create(ctx, faq)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqResponse{
		ID:       faq.ID,
		UUID:     faq.UUID,
		Question: faq.Question,
		Answer:   faq.Answer,
		Status:   faq.Status,
	}

	return result, nil
}

func (u *faq) Update(ctx context.Context, payload dto.UpdateFaqRequest) (result dto.FaqResponse, err error) {
	var (
		faq = &model.FaqModel{
			FaqEntity: model.FaqEntity{
				Question:      payload.Question,
				Answer:        payload.Answer,
				Status:        payload.Status,
				FaqCategoryID: payload.FaqCategoryID,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		faq, err = u.Repo.Faq.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.Question != "" {
			faq.Question = payload.Question
		}

		if payload.Answer != "" {
			faq.Answer = payload.Answer
		}

		if payload.Status != 0 {
			faq.Status = payload.Status
		}

		if payload.FaqCategoryID != 0 {
			faq.FaqCategoryID = payload.FaqCategoryID
		}

		_, err = u.Repo.Faq.UpdateByID(ctx, payload.ID, *faq)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqResponse{
		ID:       faq.ID,
		UUID:     faq.UUID,
		Question: faq.Question,
		Answer:   faq.Answer,
		Status:   faq.Status,
	}

	return result, nil
}

func (u *faq) Delete(ctx context.Context, payload dto.ByIDRequest) (result dto.FaqResponse, err error) {
	var data *model.FaqModel

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		data, err = u.Repo.Faq.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.Faq.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.FaqResponse{
		ID:       data.ID,
		UUID:     data.UUID,
		Question: data.Question,
		Answer:   data.Answer,
		Status:   data.Status,
	}

	return result, nil
}
