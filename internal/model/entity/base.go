package entity

import (
	"time"

	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/ctxval"
	"gorm.io/gorm"
)

type Entity struct {
	ID uint64 `json:"id" gorm:"primaryKey;autoIncrement"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if m.CreatedAt == nil {
		m.CreatedAt = &now
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = &now
	}
	return
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	if m.UpdatedAt == nil {
		m.UpdatedAt = &now
	}

	authCtx := ctxval.GetAuthValue(tx.Statement.Context)
	if authCtx != nil {
		//m.UpdatedBy = authCtx.Name
	}
	return
}

type MasterEntity struct {
	Code string `json:"code" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
}

type Tabler interface {
	TableName() string
}
