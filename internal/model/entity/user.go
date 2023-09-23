package entity

import (
	"github.com/google/uuid"
	"time"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	UUID                uuid.UUID `json:"uuid" gorm:"not null"`
	CompanyID           *uint64   `json:"company_id" gorm:"not null"`
	Name                *string   `json:"name" gorm:"not null"`
	Email               *string   `json:"email" gorm:"not null"`
	LastIPAddress       *string   `json:"-" gorm:"not null" swaggerignore:"true"`
	LastIPAddressAccess time.Time `json:"-" gorm:"not null" swaggerignore:"true"`
	IsContactPerson     *bool     `json:"is_contact_person" gorm:"not null"`
	Status              *uint8    `json:"status" gorm:"not null"`
	PasswordHash        string    `json:"-" gorm:"not null" swaggerignore:"true"`
	Password            string    `json:"-" gorm:"-" swaggerignore:"true"`

	// fk
	RoleID *uint64 `json:"role_id" gorm:"not null"`
}

type UserModel struct {
	Entity
	UserEntity

	// relations
	UserProfile *UserProfileModel `json:"-" gorm:"foreignKey:UserID"`
}

func (*UserModel) TableName() string {
	return "users"
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	err = m.Entity.BeforeCreate(tx)
	if err != nil {
		return
	}

	m.HashPassword()
	m.Password = ""
	return
}

func (m *UserModel) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserModel) GenerateToken() (string, error) {
	var (
		jwtKey = config.Config.JWT.Secret
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    m.ID,
		"role_id":    m.RoleID,
		"company_id": m.CompanyID,
		"name":       m.Name,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
