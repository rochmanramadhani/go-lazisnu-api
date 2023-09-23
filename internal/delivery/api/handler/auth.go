package handler

import (
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type (
	auth struct {
		Factory factory.Factory
	}
	Auth interface {
		Route(g *echo.Group)
		Login(c echo.Context) error
		Register(c echo.Context) error
	}
)

func NewAuth(f factory.Factory) *auth {
	return &auth{f}
}

func (h *auth) Route(g *echo.Group) {
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
}

// Login a user to get a token.
// @Summary Login user
// @Description Authenticate and log in a user.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginAuthRequest true "Request body containing user login credentials"
// @Success 200 {object} dto.AuthLoginResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/auth/login [post]
func (h *auth) Login(c echo.Context) error {
	payload := new(dto.LoginAuthRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Login(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// Register
// @Summary Register user
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterAuthRequest true "request body"
// @Success 200 {object} dto.AuthRegisterResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/auth/register [post]
func (h *auth) Register(c echo.Context) error {
	payload := new(dto.RegisterAuthRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Register(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
