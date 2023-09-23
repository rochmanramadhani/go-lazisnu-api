package dto

import (
	"github.com/google/uuid"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"html"
	"strings"
)

// request
type (
	CreateFaqCategoryRequest struct {
		Name   string `json:"name" validate:"required" example:"Category 1"`
		Status uint8  `json:"status" example:"1"`
	}

	UpdateFaqCategoryRequest struct {
		ID     uint64 `json:"-"`
		Name   string `json:"name" validate:"required" example:"Category 1"`
		Status uint8  `json:"status" example:"1"`
	}
)

func (dto *CreateFaqCategoryRequest) Prepare() {
	dto.Name = html.EscapeString(strings.TrimSpace(dto.Name))
}

func (dto *UpdateFaqCategoryRequest) Prepare() {
	dto.Name = html.EscapeString(strings.TrimSpace(dto.Name))
}

// response
type (
	FaqCategoryResponse struct {
		ID     uint64        `json:"id" example:"1"`
		UUID   uuid.UUID     `json:"uuid" example:"c954ea37-91ea-4f2f-837d-00a84d4af106"`
		Name   string        `json:"name" example:"Category 1"`
		Status uint8         `json:"status" example:"1"`
		Faqs   []FaqResponse `json:"faqs,omitempty"`
	}
	FaqCategoryResponseDoc struct {
		Meta res.Meta            `json:"meta"`
		Data FaqCategoryResponse `json:"data"`
	}
	FaqCategoryResponseListDoc struct {
		Meta res.Meta              `json:"meta"`
		Data []FaqCategoryResponse `json:"data"`
	}
)
