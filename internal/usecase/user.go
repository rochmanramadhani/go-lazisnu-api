package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/ctxval"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/trxmanager"
)

type User interface {
	Find(ctx context.Context, filterParam abstraction.Filter) ([]dto.UserResponse, abstraction.PaginationInfo, error)
	FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.UserResponse, error)
	Create(ctx context.Context, payload dto.CreateUserRequest) (dto.UserResponse, error)
	Update(ctx context.Context, payload dto.UpdateUserRequest) (dto.UserResponse, error)
	Delete(ctx context.Context, payload dto.ByIDRequest) (dto.UserResponse, error)
}

type user struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewUser(cfg *config.Configuration, f repository.Factory) User {
	return &user{cfg, f}
}

func (u *user) Find(ctx context.Context, filterParam abstraction.Filter) (result []dto.UserResponse, pagination abstraction.PaginationInfo, err error) {
	var search *abstraction.Search
	if filterParam.Search != "" {
		searchQuery := "lower(name) LIKE ? OR lower(email) Like ?"
		searchVal := "%" + strings.ToLower(filterParam.Search) + "%"
		search = &abstraction.Search{
			Query: searchQuery,
			Args:  []interface{}{searchVal, searchVal},
		}
	}

	users, info, err := u.Repo.User.Find(ctx, filterParam, search)
	if err != nil {
		return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	pagination = *info

	for _, user := range users {
		role, err := u.Repo.Role.FindByID(ctx, *user.RoleID)
		if err != nil {
			return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
		}

		userProfile, err := u.Repo.UserProfile.FindByUserID(ctx, user.ID)
		if err != nil {
			return nil, pagination, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
		}

		result = append(result, dto.UserResponse{
			ID:                  user.ID,
			UUID:                user.UUID,
			CompanyID:           *user.CompanyID,
			RoleID:              0, // TODO: do i need to show this?
			Name:                *user.Name,
			Email:               *user.Email,
			LastIPAddress:       *user.LastIPAddress,
			LastIPAddressAccess: user.LastIPAddressAccess,
			IsContactPerson:     *user.IsContactPerson,
			Status:              *user.Status,
			Role: &dto.RoleResponse{
				ID:          role.ID,
				UUID:        role.UUID,
				Name:        role.Name,
				Description: role.Description,
			},
			UserProfile: &dto.UserProfileResponse{
				Address:  *userProfile.Address,
				Phone:    *userProfile.Phone,
				FilePath: userProfile.FilePath,
			},
		})
	}

	return result, pagination, nil
}

func (u *user) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.UserResponse, error) {
	var result dto.UserResponse

	user, err := u.Repo.User.FindByID(ctx, payload.ID)
	if err != nil {
		return result, err
	}

	role, err := u.Repo.Role.FindByID(ctx, *user.RoleID)
	if err != nil {
		return result, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	userProfile, err := u.Repo.UserProfile.FindByUserID(ctx, user.ID)
	if err != nil {
		return result, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	result = dto.UserResponse{
		ID:                  user.ID,
		UUID:                user.UUID,
		CompanyID:           *user.CompanyID,
		RoleID:              0, // TODO: do i need to show this?
		Name:                *user.Name,
		Email:               *user.Email,
		LastIPAddress:       *user.LastIPAddress,
		LastIPAddressAccess: user.LastIPAddressAccess,
		IsContactPerson:     *user.IsContactPerson,
		Status:              *user.Status,
		Role: &dto.RoleResponse{
			ID:          role.ID,
			UUID:        role.UUID,
			Name:        role.Name,
			Description: role.Description,
		},
		UserProfile: &dto.UserProfileResponse{
			Address:  *userProfile.Address,
			Phone:    *userProfile.Phone,
			FilePath: userProfile.FilePath,
		},
	}

	return result, nil
}

func (u *user) Create(ctx context.Context, payload dto.CreateUserRequest) (result dto.UserResponse, err error) {
	var (
		uid = uuid.New()

		user = model.UserModel{
			UserEntity: model.UserEntity{
				UUID:                uid,
				CompanyID:           &payload.CompanyID,
				Name:                &payload.Name,
				Email:               &payload.Email,
				LastIPAddress:       nil,
				LastIPAddressAccess: time.Now(),
				IsContactPerson:     &payload.IsContactPerson,
				Status:              &payload.Status,
				Password:            payload.Password,
			},
		}
		userProfile = &model.UserProfileModel{
			UserProfileEntity: model.UserProfileEntity{
				Address: &payload.Address,
				Phone:   &payload.Phone,
			},
		}
	)

	if payload.Password != "" {
		user.HashPassword()
		user.Password = ""
	}

	var role *model.RoleModel
	if payload.RoleName != nil {
		role, err = u.Repo.Role.FindByName(ctx, *payload.RoleName)
		if err != nil {
			return result, res.ErrorBuilder(res.Constant.Error.BadRequest, err, "role name invalid")
		}
	} else {
		role, err = u.Repo.Role.FindByID(ctx, *payload.RoleID)
		if err != nil {
			return result, res.ErrorBuilder(res.Constant.Error.BadRequest, err, "role id invalid")
		}
	}
	user.RoleID = &role.ID

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		createdUser, err := u.Repo.User.Create(ctx, user)
		if err != nil {
			return err
		}
		userProfile.UserID = createdUser.ID

		_, err = u.Repo.UserProfile.Create(ctx, *userProfile)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		ID:                  user.ID,
		UUID:                user.UUID,
		CompanyID:           *user.CompanyID,
		RoleID:              0, // TODO: do i need to show this?
		Name:                *user.Name,
		Email:               *user.Email,
		LastIPAddress:       *user.LastIPAddress,
		LastIPAddressAccess: user.LastIPAddressAccess,
		IsContactPerson:     *user.IsContactPerson,
		Status:              *user.Status,
		Role: &dto.RoleResponse{
			ID:          role.ID,
			UUID:        role.UUID,
			Name:        role.Name,
			Description: role.Description,
		},
		UserProfile: &dto.UserProfileResponse{
			Address:  *userProfile.Address,
			Phone:    *userProfile.Phone,
			FilePath: userProfile.FilePath,
		},
	}

	return result, nil
}

func (u *user) Update(ctx context.Context, payload dto.UpdateUserRequest) (result dto.UserResponse, err error) {
	fileValues := ctxval.GetUploadFileValue(ctx)
	files := *fileValues

	var (
		user = &model.UserModel{
			UserEntity: model.UserEntity{
				LastIPAddress:       nil, // TODO: get ip address
				LastIPAddressAccess: time.Now(),
				IsContactPerson:     payload.IsContactPerson,
				Status:              payload.Status,
				PasswordHash:        payload.Password,
				Password:            payload.Password,
				RoleID:              payload.RoleID,
			},
		}
		userProfile = &model.UserProfileModel{
			UserProfileEntity: model.UserProfileEntity{
				Address:  payload.Address,
				Phone:    payload.Phone,
				FilePath: files[0].FilePath,
				UserID:   payload.ID,
			},
		}
	)

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		user, err = u.Repo.User.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.IsContactPerson != nil {
			user.IsContactPerson = payload.IsContactPerson
		}

		if payload.Status != nil {
			user.Status = payload.Status
		}

		if payload.Password != "" {
			user.HashPassword()
			user.Password = ""
		}

		if payload.RoleID != nil {
			user.RoleID = payload.RoleID
		}

		_, err = u.Repo.User.UpdateByID(ctx, payload.ID, *user)
		if err != nil {
			return err
		}

		userProfile, err = u.Repo.UserProfile.FindByUserID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}

		if payload.Address != nil {
			userProfile.Address = payload.Address
		}

		if payload.Phone != nil {
			userProfile.Phone = payload.Phone
		}

		if files[0].FilePath != "" {
			filePath := userProfile.FilePath
			if err := os.Remove(filePath); err != nil {
				logrus.Error(err)
				return res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
			}
			userProfile.FilePath = files[0].FilePath
		}

		_, err = u.Repo.UserProfile.UpdateByUserID(ctx, payload.ID, *userProfile)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		ID:                  user.ID,
		UUID:                user.UUID,
		CompanyID:           *user.CompanyID,
		RoleID:              *user.RoleID,
		Name:                *user.Name,
		Email:               *user.Email,
		LastIPAddress:       *user.LastIPAddress,
		LastIPAddressAccess: user.LastIPAddressAccess,
		IsContactPerson:     *user.IsContactPerson,
		Status:              *user.Status,
		Role:                nil,
		UserProfile: &dto.UserProfileResponse{
			Address:  *userProfile.Address,
			Phone:    *userProfile.Phone,
			FilePath: userProfile.FilePath,
		},
	}

	return result, nil
}

func (u *user) Delete(ctx context.Context, payload dto.ByIDRequest) (result dto.UserResponse, err error) {
	var user *model.UserModel

	if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
		user, err = u.Repo.User.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.UserProfile.DeleteByUserID(ctx, payload.ID)
		if err != nil {
			return err
		}

		err = u.Repo.User.DeleteByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.UserResponse{
		ID:                  user.ID,
		UUID:                user.UUID,
		CompanyID:           *user.CompanyID,
		RoleID:              *user.RoleID,
		Name:                *user.Name,
		Email:               *user.Email,
		LastIPAddress:       *user.LastIPAddress,
		LastIPAddressAccess: user.LastIPAddressAccess,
		IsContactPerson:     *user.IsContactPerson,
		Status:              *user.Status,
		Role:                nil,
		UserProfile:         nil,
	}

	return result, nil
}
