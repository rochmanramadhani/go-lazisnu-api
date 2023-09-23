package entity

import (
	"github.com/google/uuid"
)

type DonationTypeEntity struct {
	UUID uuid.UUID `json:"uuid" gorm:"not null"`
	Name string    `json:"name" gorm:"not null"`
}

type DonationTypeModel struct {
	Entity
	DonationTypeEntity
}

func (DonationTypeModel) TableName() string {
	return "donation_types"
}
