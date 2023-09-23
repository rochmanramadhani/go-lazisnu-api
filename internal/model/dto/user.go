package dto

import (
	"github.com/google/uuid"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"time"
)

// request
type (
	CreateUserRequest struct {
		CompanyID       uint64  `json:"company_id" validate:"required" example:"1"`
		Name            string  `json:"name" validate:"required" example:"Admin User"`
		Email           string  `json:"email" validate:"required,email" example:"admin@example.com"`
		IsContactPerson bool    `json:"is_contact_person" validate:"required" example:"true"`
		Status          uint8   `json:"status" validate:"required" example:"1"`
		Password        string  `json:"password" validate:"required" example:"password"`
		RoleID          *uint64 `json:"role_id,omitempty" example:"1"`
		RoleName        *string `json:"role_name,omitempty" example:"admin"`
		Address         string  `json:"address,omitempty" example:"Jl. Jend. Sudirman No. 1"`
		Phone           string  `json:"phone,omitempty" example:"08123456789"`
	}

	UpdateUserRequest struct {
		ID              uint64  `json:"-"`
		IsContactPerson *bool   `form:"is_contact_person" json:"is_contact_person"`
		Password        string  `form:"password" json:"password"`
		Status          *uint8  `form:"status" json:"status"`
		RoleID          *uint64 `form:"role_id" json:"role_id"`
		RoleName        *string `form:"role_name" json:"role_name"`
		Address         *string `form:"address" json:"address"`
		Phone           *string `form:"phone" json:"phone"`
	}
)

// response
type (
	UserResponse struct {
		ID                  uint64    `json:"id" example:"1"`
		UUID                uuid.UUID `json:"uuid" example:"c954ea37-91ea-4f2f-837d-00a84d4af106"`
		CompanyID           uint64    `json:"company_id" example:"1"`
		RoleID              uint64    `json:"role_id,omitempty" example:"1"`
		Name                string    `json:"name" example:"Admin User"`
		Email               string    `json:"email" example:"admin@example.com"`
		LastIPAddress       string    `json:"last_ip_address" example:"127.0.0.1"`
		LastIPAddressAccess time.Time `json:"last_ip_address_access" example:"2022-01-01T00:00:00Z"`
		IsContactPerson     bool      `json:"is_contact_person" example:"true"`
		Status              uint8     `json:"status" example:"1"`

		Role        *RoleResponse        `json:"role,omitempty"`
		UserProfile *UserProfileResponse `json:"user_profile,omitempty"`
	}
	UserProfileResponse struct {
		Address  string `json:"address" example:"Jl. Jend. Sudirman No. 1"`
		Phone    string `json:"phone" example:"08123456789"`
		FilePath string `json:"file_path" example:"/path/to/file.jpg"`
	}
	UserResponseDoc struct {
		Meta res.Meta     `json:"meta"`
		Data UserResponse `json:"data"`
	}
	UserResponseListDoc struct {
		Meta res.Meta       `json:"meta"`
		Data []UserResponse `json:"data"`
	}
)
