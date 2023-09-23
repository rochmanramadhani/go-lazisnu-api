package dto

import (
	"github.com/google/uuid"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
)

// request
type (
	CreateRoleRequest struct {
		Name        string `json:"name" validate:"required" example:"admin"`
		Description string `json:"description" validate:"required" example:"Administrator"`
	}

	UpdateRoleRequest struct {
		ID          uint64 `json:"-"`
		Name        string `json:"name" validate:"required" example:"admin"`
		Description string `json:"description" example:"Administrator"`
	}
)

// response
type (
	RoleResponse struct {
		ID          uint64    `json:"id" example:"1"`
		UUID        uuid.UUID `json:"uuid" example:"c954ea37-91ea-4f2f-837d-00a84d4af106"`
		Name        string    `json:"name" example:"admin"`
		Description string    `json:"description" example:"Administrator"`
	}
	RoleResponseDoc struct {
		Meta res.Meta     `json:"meta"`
		Data RoleResponse `json:"data"`
	}
	RoleResponseListDoc struct {
		Meta res.Meta       `json:"meta"`
		Data []RoleResponse `json:"data"`
	}
)
