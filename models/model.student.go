package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar;not null"`
	Npm  uint64 `json:"npm" gorm:"type:bigint;not null"`
	Fak  string `json:"fak" gorm:"type:varchar;not null"`
	Bid  string `json:"bid" gorm:"type:varchar;not null"`
	// Teachers  pq.StringArray `json:"teachers" gorm:"type:text[];constraint:onupdate:cascade,ondelete:cascade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (entity *Student) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Student) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
