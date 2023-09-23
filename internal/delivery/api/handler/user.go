package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api/middleware"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"strconv"
)

type (
	user struct {
		Factory factory.Factory
	}
	User interface {
		Route(g *echo.Group)
		Get(c echo.Context) error
		GetByID(c echo.Context) error
		Create(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)

func NewUser(f factory.Factory) User {
	return &user{f}
}

func (h *user) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication, middleware.UploadFiles("user"))
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get user
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @param request query abstraction.Filter true "request query"
// @Param entity query model.UserEntity false "entity query"
// @Success 200 {object} dto.UserResponseListDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/users [get]
func (h *user) Get(c echo.Context) error {
	filter := abstraction.NewFilterBuiler[model.UserEntity](c, "users")
	if err := c.Bind(filter.Payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	filter.Bind()

	result, pagination, err := h.Factory.Usecase.User.Find(c.Request().Context(), *filter.Payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get users success", &pagination).Send(c)
}

// GetByID user by id
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/users/{id} [get]
func (h *user) GetByID(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.User.FindByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create user
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/users [post]
func (h *user) Create(c echo.Context) error {
	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Create(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update user
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param files formData file true "files"
// @Param request formData dto.UpdateUserRequest false "request body"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/users/{id} [put]
func (h *user) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	payload := new(dto.UpdateUserRequest)
	payload.ID = id
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Update(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete user
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.UserResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/users/{id} [delete]
func (h *user) Delete(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.User.Delete(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
