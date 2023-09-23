package dto

import (
	"github.com/google/uuid"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
)

// request
type (
	CreateDonationTypeRequest struct {
		Name string `json:"name" validate:"required" example:"admin"`
	}

	UpdateDonationTypeRequest struct {
		ID   uint64 `json:"-"`
		Name string `json:"name" validate:"required" example:"admin"`
	}
)

// response
type (
	DonationTypeResponse struct {
		ID   uint64    `json:"id" example:"1"`
		UUID uuid.UUID `json:"uuid" example:"c954ea37-91ea-4f2f-837d-00a84d4af106"`
		Name string    `json:"name" example:"admin"`
	}
	DonationTypeResponseDoc struct {
		Meta res.Meta             `json:"meta"`
		Data DonationTypeResponse `json:"data"`
	}
	DonationTypeResponseListDoc struct {
		Meta res.Meta               `json:"meta"`
		Data []DonationTypeResponse `json:"data"`
	}
)
