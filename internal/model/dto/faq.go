package dto

import (
	"github.com/google/uuid"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"html"
	"strings"
)

// request
type (
	CreateFaqRequest struct {
		FaqCategoryID uint64 `json:"faq_category_id" validate:"required" example:"1"`
		Question      string `json:"question" validate:"required" example:"Question 1"`
		Answer        string `json:"answer" validate:"required" example:"Answer 1"`
		Status        uint8  `json:"status" example:"1"`
	}

	UpdateFaqRequest struct {
		ID            uint64 `json:"-"`
		FaqCategoryID uint64 `json:"faq_category_id" validate:"required" example:"1"`
		Question      string `json:"question" validate:"required" example:"Question 1"`
		Answer        string `json:"answer" validate:"required" example:"Answer 1"`
		Status        uint8  `json:"status" example:"1"`
	}
)

func (dto *CreateFaqRequest) Prepare() {
	dto.Question = html.EscapeString(strings.TrimSpace(dto.Question))
	dto.Answer = html.EscapeString(strings.TrimSpace(dto.Answer))
}

func (dto *UpdateFaqRequest) Prepare() {
	dto.Question = html.EscapeString(strings.TrimSpace(dto.Question))
	dto.Answer = html.EscapeString(strings.TrimSpace(dto.Answer))
}

// response
type (
	FaqResponse struct {
		ID       uint64    `json:"id" example:"1"`
		UUID     uuid.UUID `json:"uuid" example:"c954ea37-91ea-4f2f-837d-00a84d4af106"`
		Question string    `json:"question" example:"Question 1"`
		Answer   string    `json:"answer" example:"Answer 1"`
		Status   uint8     `json:"status" example:"1"`
	}
	FaqResponseDoc struct {
		Meta res.Meta    `json:"meta"`
		Data FaqResponse `json:"data"`
	}
	FaqResponseListDoc struct {
		Meta res.Meta      `json:"meta"`
		Data []FaqResponse `json:"data"`
	}
)
