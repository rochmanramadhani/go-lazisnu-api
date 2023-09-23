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
	faqCategory struct {
		Factory factory.Factory
	}
	FaqCategory interface {
		Route(g *echo.Group)
		Get(c echo.Context) error
		GetByID(c echo.Context) error
		Create(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)

func NewFaqCategory(f factory.Factory) FaqCategory {
	return &faqCategory{f}
}

func (h *faqCategory) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.DELETE("/:id", h.Delete, middleware.Authentication)
}

// Get faq category
// @Summary Get faq category
// @Description Get faq category
// @Tags faq-categories
// @Accept json
// @Produce json
// @param request query abstraction.Filter true "request query"
// @Param entity query model.FaqCategoryEntity false "entity query"
// @Success 200 {object} dto.FaqCategoryResponseListDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faq-categories [get]
func (h *faqCategory) Get(c echo.Context) error {
	filter := abstraction.NewFilterBuiler[model.FaqCategoryEntity](c, "faq_categories")
	if err := c.Bind(filter.Payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	filter.Bind()

	result, pagination, err := h.Factory.Usecase.FaqCategory.Find(c.Request().Context(), *filter.Payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get faq categories success", &pagination).Send(c)
}

// GetByID faq category by id
// @Summary Get faq category by id
// @Description Get faq category by id
// @Tags faq-categories
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.FaqCategoryResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faq-categories/{id} [get]
func (h *faqCategory) GetByID(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	result, err := h.Factory.Usecase.FaqCategory.FindByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(result).Send(c)
}

// Create faq category
// @Summary Create faq category
// @Description Create faq category
// @Tags faq-categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateFaqCategoryRequest true "request body"
// @Success 200 {object} dto.FaqCategoryResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faq-categories [post]
func (h *faqCategory) Create(c echo.Context) error {
	payload := new(dto.CreateFaqCategoryRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	payload.Prepare()
	result, err := h.Factory.Usecase.FaqCategory.Create(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Update faq category
// @Summary Update faq category
// @Description Update faq category
// @Tags faq-categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Param request body dto.UpdateFaqCategoryRequest true "request body"
// @Success 200 {object} dto.FaqCategoryResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faq-categories/{id} [put]
func (h *faqCategory) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	payload := new(dto.UpdateFaqCategoryRequest)
	payload.ID = id
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	payload.Prepare()
	result, err := h.Factory.Usecase.FaqCategory.Update(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete faq category
// @Summary Delete faq category
// @Description Delete faq category
// @Tags faq-categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id path"
// @Success 200 {object} dto.FaqCategoryResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/v1/faq-categories/{id} [delete]
func (h *faqCategory) Delete(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.FaqCategory.Delete(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
