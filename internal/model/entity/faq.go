package entity

import (
	"github.com/google/uuid"
)

type FaqEntity struct {
	UUID     uuid.UUID `json:"uuid" gorm:"not null"`
	Question string    `json:"question" gorm:"not null"`
	Answer   string    `json:"answer" gorm:"not null"`
	Status   uint8     `json:"status" gorm:"not null"`

	// fk
	FaqCategoryID uint64 `json:"faq_category_id" gorm:"not null"`
}

type FaqModel struct {
	Entity
	FaqEntity
}

func (FaqModel) TableName() string {
	return "faqs"
}
