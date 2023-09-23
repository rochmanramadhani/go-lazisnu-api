package entity

import (
	"github.com/google/uuid"
)

type FaqCategoryEntity struct {
	UUID   uuid.UUID `json:"uuid" gorm:"not null"`
	Name   string    `json:"name" gorm:"not null"`
	Status uint8     `json:"status" gorm:"not null"`
}

type FaqCategoryModel struct {
	Entity
	FaqCategoryEntity

	// relations
	Faqs []FaqModel `json:"faqs" gorm:"foreignKey:FaqCategoryID"`
}

func (FaqCategoryModel) TableName() string {
	return "faq_categories"
}
