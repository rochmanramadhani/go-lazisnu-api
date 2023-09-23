package handler

import (
	"strconv"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api/middleware"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type (
	donationType struct {
		Factory factory.Factory
	}
	DonationType interface {
		Route(g *echo.Group)
		Get(c echo.Context) error
		GetByID(c echo.Context) error
		Create(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)

func NewDonationType(f factory.Factory) DonationType {
	return &donationType{f}
}

func (h *donationType) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get donation type
// @Summary Get donation type
// @Description Get donation type
// @Tags donation-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @param request query abstraction.Filter true "request query"
// @Param entity query model.DonationTypeEntity false "entity query"
// @Success 200 {object} dto.DonationTypeResponseListDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/donation-types [get]
func (h *donationType) Get(c echo.Context) error {
	filter := abstraction.NewFilterBuiler[model.DonationTypeEntity](c, "donation_types")
	if err := c.Bind(filter.Payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	filter.Bind()

	result, pagination, err := h.Factory.Usecase.DonationType.Find(c.Request().Context(), *filter.Payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get donation types success", &pagination).Send(c)
}

// GetByID donation type by id
// @Summary Get donation type by id
// @Description Get donation type by id
// @Tags donation-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.DonationTypeResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/donation-types/{id} [get]
func (h *donationType) GetByID(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.DonationType.FindByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create donation type
// @Summary Create donation type
// @Description Create donation type
// @Tags donation-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateDonationTypeRequest true "request body"
// @Success 200 {object} dto.DonationTypeResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/donation-types [post]
func (h *donationType) Create(c echo.Context) error {
	payload := new(dto.CreateDonationTypeRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.DonationType.Create(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update donation type
// @Summary Update donation type
// @Description Update donation type
// @Tags donation-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param request body dto.UpdateDonationTypeRequest true "request body"
// @Success 200 {object} dto.DonationTypeResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/donation-types/{id} [put]
func (h *donationType) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	payload := new(dto.UpdateDonationTypeRequest)
	payload.ID = id
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.DonationType.Update(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete donation type
// @Summary Delete donation type
// @Description Delete donation type
// @Tags donation-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.DonationTypeResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/donation-types/{id} [delete]
func (h *donationType) Delete(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.DonationType.Delete(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
