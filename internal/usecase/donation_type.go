package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/trxmanager"
)

type DonationType interface {
	Find(ctx context.Context, filterParam abstraction.Filter) ([]dto.DonationTypeResponse, abstraction.PaginationInfo, error)
	FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.DonationTypeResponse, error)
	Create(ctx context.Context, payload dto.CreateDonationTypeRequest) (dto.DonationTypeResponse, error)
	Update(ctx context.Context, payload dto.UpdateDonationTypeRequest) (dto.DonationTypeResponse, error)
	Delete(ctx context.Context, payload dto.ByIDRequest) (dto.DonationTypeResponse, error)
}

type donationType struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewDonationType(cfg *config.Configuration, f repository.Factory) DonationType {
	return &donationType{cfg, f}
}

func (u *donationType) Find(ctx context.Context, filterParam abstraction.Filter) (result []dto.DonationTypeResponse, pagination abstraction.PaginationInfo, err error) {
	var search *abstraction.Search
	if filterParam.Search != "" {
		searchQuery := "lower(code) LIKE ? OR lower(name) LIKE ?"
		searchVal := "%" + strings.ToLower(filterParam.Search) + "%"
		search = &abstraction.Search{
			Query: searchQuery,
			Args:  []interface{}{searchVal, searchVal},
		}
	}

	donationTypes, info, err := u.Repo.DonationType.Find(ctx, filterParam, search)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	pagination = *info

	for _, donationType := range donationTypes {
		result = append(result, dto.DonationTypeResponse{
			ID:   donationType.ID,
			UUID: donationType.UUID,
			Name: donationType.Name,
		})
	}

	return result, pagination, nil
}

func (u *donationType) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.DonationTypeResponse, error) {
	var result dto.DonationTypeResponse

	donationType, err := u.Repo.DonationType.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	result = dto.DonationTypeResponse{
		ID:   donationType.ID,
		UUID: donationType.UUID,
		Name: donationType.Name,
	}

	return result, nil
}

func (u *donationType) Create(ctx context.Context, payload dto.CreateDonationTypeRequest) (result dto.DonationTypeResponse, err error) {
	var (
		donationType = model.DonationTypeModel{
			DonationTypeEntity: model.DonationTypeEntity{
				UUID: uuid.New(),
				Name: payload.Name,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		donationType, err = u.Repo.DonationType.Create(ctx, donationType)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.DonationTypeResponse{
		ID:   donationType.ID,
		UUID: donationType.UUID,
		Name: donationType.Name,
	}

	return result, nil
}

func (u *donationType) Update(ctx context.Context, payload dto.UpdateDonationTypeRequest) (result dto.DonationTypeResponse, err error) {
	var (
		donationType = &model.DonationTypeModel{
			DonationTypeEntity: model.DonationTypeEntity{
				Name: payload.Name,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		donationType, err = u.Repo.DonationType.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.Name != "" {
			donationType.Name = payload.Name
		}

		_, err = u.Repo.DonationType.UpdateByID(ctx, payload.ID, *donationType)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.DonationTypeResponse{
		ID:   donationType.ID,
		UUID: donationType.UUID,
		Name: donationType.Name,
	}

	return result, nil
}

func (u *donationType) Delete(ctx context.Context, payload dto.ByIDRequest) (result dto.DonationTypeResponse, err error) {
	var data *model.DonationTypeModel

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		data, err = u.Repo.DonationType.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.DonationType.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.DonationTypeResponse{
		ID:   data.ID,
		UUID: data.UUID,
		Name: data.Name,
	}

	return result, nil
}
