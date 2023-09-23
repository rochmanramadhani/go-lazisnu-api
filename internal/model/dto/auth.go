package dto

import (
	"github.com/google/uuid"
	model "github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	res "github.com/rochmanramadhani/go-lazisnu-api/pkg/util/response"
	"strings"
	"time"
)

// request
type (
	LoginAuthRequest struct {
		Email    string `json:"email" validate:"required,email" example:"admin@example.com"`
		Password string `json:"password" validate:"required,min=8" example:"password"`
	}
	RegisterAuthRequest struct {
		model.UserEntity
	}
)

func (p *LoginAuthRequest) Prepare() {
	p.Email = strings.TrimSpace(p.Email)
	p.Password = strings.TrimSpace(p.Password)
}

// response
type (
	AuthLoginResponse struct {
		Token               string    `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjb21wYW55X2lkIjo1OSwiZXhwIjoxNjk1MjgwNzI2LCJ1c2VyX2lkIjo1Nn0.lH57qsT3eGj-hVxk6LdGwQPo_qC3Ep0j1_eC9qAXJi8"`
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

		Role *RoleResponse `json:"role,omitempty"`
	}
	AuthLoginResponseDoc struct {
		Meta res.Meta          `json:"meta"`
		Data AuthLoginResponse `json:"data"`
	}

	AuthRegisterResponse struct {
		model.UserModel
	}
	AuthRegisterResponseDoc struct {
		Meta res.Meta             `json:"meta"`
		Data AuthRegisterResponse `json:"data"`
	}
)
