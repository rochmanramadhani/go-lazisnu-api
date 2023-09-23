package entity

import (
	"github.com/google/uuid"
)

type RoleEntity struct {
	UUID        uuid.UUID `json:"uuid" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
}

type RoleModel struct {
	Entity
	RoleEntity

	// relations
	Users *[]UserModel `json:"-" gorm:"foreignKey:RoleID"`
}

func (RoleModel) TableName() string {
	return "roles"
}
