package handler

import (
	"github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api/middleware"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/dto"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	faq struct {
		Factory factory.Factory
	}
	Faq interface {
		Route(g *echo.Group)
		Get(c echo.Context) error
		GetByID(c echo.Context) error
		Create(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)

func NewFaq(f factory.Factory) Faq {
	return &faq{f}
}

func (h *faq) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get faq
// @Summary Get faq
// @Description Get faq
// @Tags faq
// @Accept json
// @Produce json
// @param request query abstraction.Filter true "request query"
// @Param entity query model.FaqEntity false "entity query"
// @Success 200 {object} dto.FaqResponseListDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faqs [get]
func (h *faq) Get(c echo.Context) error {
	filter := abstraction.NewFilterBuiler[model.FaqEntity](c, "faqs")
	if err := c.Bind(filter.Payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	filter.Bind()

	result, pagination, err := h.Factory.Usecase.Faq.Find(c.Request().Context(), *filter.Payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get faqs success", &pagination).Send(c)
}

// GetByID faq by id
// @Summary Get faq by id
// @Description Get faq by id
// @Tags faq
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.FaqResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faqs/{id} [get]
func (h *faq) GetByID(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.Faq.FindByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create faq
// @Summary Create faq
// @Description Create faq
// @Tags faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateFaqRequest true "request body"
// @Success 200 {object} dto.FaqResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faqs [post]
func (h *faq) Create(c echo.Context) error {
	payload := new(dto.CreateFaqRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	payload.Prepare()
	result, err := h.Factory.Usecase.Faq.Create(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update faq
// @Summary Update faq
// @Description Update faq
// @Tags faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param request body dto.UpdateFaqRequest true "request body"
// @Success 200 {object} dto.FaqResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faqs/{id} [put]
func (h *faq) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	payload := new(dto.UpdateFaqRequest)
	payload.ID = id
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	payload.Prepare()
	result, err := h.Factory.Usecase.Faq.Update(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete faq
// @Summary Delete faq
// @Description Delete faq
// @Tags faq
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.FaqResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faqs/{id} [delete]
func (h *faq) Delete(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Faq.Delete(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
