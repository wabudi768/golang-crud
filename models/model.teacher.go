package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar;not null"`
	Matkul    string    `json:"matkul" gorm:"type:varchar;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (entity *Teacher) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Teacher) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
