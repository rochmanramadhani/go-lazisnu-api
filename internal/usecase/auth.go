package usecase

import (
	"context"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	resConstant "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response/constant"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Auth interface {
	Login(ctx context.Context, payload dto.LoginAuthRequest) (dto.AuthLoginResponse, error)
	Register(ctx context.Context, payload dto.RegisterAuthRequest) (dto.AuthRegisterResponse, error)
}

type auth struct {
	Cfg  *config.Configuration
	Repo repository.Factory
}

func NewAuth(cfg *config.Configuration, f repository.Factory) Auth {
	return &auth{cfg, f}
}

func (u *auth) Login(ctx context.Context, payload dto.LoginAuthRequest) (dto.AuthLoginResponse, error) {
	var result dto.AuthLoginResponse

	user, err := u.Repo.User.FindByEmail(ctx, payload.Email)
	if user == nil {
		return result, res.CustomErrorBuilder(http.StatusUnauthorized, resConstant.E_INVALID_CREDENTIALS, err, "")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.CustomErrorBuilder(http.StatusUnauthorized, resConstant.E_INVALID_CREDENTIALS, err, "")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	role, err := u.Repo.Role.FindByID(ctx, *user.RoleID)
	if err != nil {
		return result, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	result = dto.AuthLoginResponse{
		Token:               token,
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
		Role: &dto.RoleResponse{
			ID:          role.ID,
			UUID:        role.UUID,
			Name:        role.Name,
			Description: role.Description,
		},
	}

	return result, nil
}

func (u *auth) Register(ctx context.Context, payload dto.RegisterAuthRequest) (dto.AuthRegisterResponse, error) {
	//var result dto.AuthRegisterResponse
	//var user model.UserModel
	//var err error
	//
	//if err = trxmanager.New(u.Repo.Db).WithTrx(ctx, func(ctx context.Context) error {
	//	userID := uuid.New().String()
	//	user = model.UserModel{
	//		Entity: model.Entity{
	//			ID: userID,
	//		},
	//		UserEntity: payload.UserEntity,
	//	}
	//	userProfile := model.UserProfileModel{
	//		UserProfileEntity: model.UserProfileEntity{
	//			Address: "xxx",
	//			Phone:   "021",
	//			UserID:  userID,
	//		},
	//	}
	//
	//	role, err := u.Repo.Role.FindByName(ctx, "customer")
	//	if err != nil {
	//		return err
	//	}
	//	user.RoleID = role.ID
	//
	//	_, err = u.Repo.User.Create(ctx, user)
	//	if err != nil {
	//		return err
	//	}
	//
	//	_, err = u.Repo.UserProfile.Create(ctx, userProfile)
	//	if err != nil {
	//		return err
	//	}
	//
	//	return nil
	//}); err != nil {
	//	return result, err
	//}
	//
	//result = dto.AuthRegisterResponse{
	//	UserModel: user,
	//}
	//
	//return result, nil
	panic("implement me")
}
