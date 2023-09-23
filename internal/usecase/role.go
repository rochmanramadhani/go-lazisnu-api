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

type Role interface {
	Find(ctx context.Context, filterParam abstraction.Filter) ([]dto.RoleResponse, abstraction.PaginationInfo, error)
	FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.RoleResponse, error)
	Create(ctx context.Context, payload dto.CreateRoleRequest) (dto.RoleResponse, error)
	Update(ctx context.Context, payload dto.UpdateRoleRequest) (dto.RoleResponse, error)
	Delete(ctx context.Context, payload dto.ByIDRequest) (dto.RoleResponse, error)
}

type role struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewRole(cfg *config.Configuration, f repository.Factory) Role {
	return &role{cfg, f}
}

func (u *role) Find(ctx context.Context, filterParam abstraction.Filter) (result []dto.RoleResponse, pagination abstraction.PaginationInfo, err error) {
	var search *abstraction.Search
	if filterParam.Search != "" {
		searchQuery := "lower(code) LIKE ? OR lower(name) LIKE ?"
		searchVal := "%" + strings.ToLower(filterParam.Search) + "%"
		search = &abstraction.Search{
			Query: searchQuery,
			Args:  []interface{}{searchVal, searchVal},
		}
	}

	roles, info, err := u.Repo.Role.Find(ctx, filterParam, search)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	pagination = *info

	for _, role := range roles {
		result = append(result, dto.RoleResponse{
			ID:          role.ID,
			UUID:        role.UUID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return result, pagination, nil
}

func (u *role) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.RoleResponse, error) {
	var result dto.RoleResponse

	role, err := u.Repo.Role.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	result = dto.RoleResponse{
		ID:          role.ID,
		UUID:        role.UUID,
		Name:        role.Name,
		Description: role.Description,
	}

	return result, nil
}

func (u *role) Create(ctx context.Context, payload dto.CreateRoleRequest) (result dto.RoleResponse, err error) {
	var (
		role = model.RoleModel{
			RoleEntity: model.RoleEntity{
				UUID:        uuid.New(),
				Name:        payload.Name,
				Description: payload.Description,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		role, err = u.Repo.Role.Create(ctx, role)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.RoleResponse{
		ID:          role.ID,
		UUID:        role.UUID,
		Name:        role.Name,
		Description: role.Description,
	}

	return result, nil
}

func (u *role) Update(ctx context.Context, payload dto.UpdateRoleRequest) (result dto.RoleResponse, err error) {
	var (
		role = &model.RoleModel{
			RoleEntity: model.RoleEntity{
				Name: payload.Name,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		role, err = u.Repo.Role.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.Name != "" {
			role.Name = payload.Name
		}

		if payload.Description != "" {
			role.Description = payload.Description
		}

		_, err = u.Repo.Role.UpdateByID(ctx, payload.ID, *role)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.RoleResponse{
		ID:          role.ID,
		UUID:        role.UUID,
		Name:        role.Name,
		Description: role.Description,
	}

	return result, nil
}

func (u *role) Delete(ctx context.Context, payload dto.ByIDRequest) (result dto.RoleResponse, err error) {
	var data *model.RoleModel

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		data, err = u.Repo.Role.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.Role.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.RoleResponse{
		ID:          data.ID,
		UUID:        data.UUID,
		Name:        data.Name,
		Description: data.Description,
	}

	return result, nil
}
