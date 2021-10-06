package models

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name   string `gorm:"type:varchar;not null"`
	Matkul string `gorm:"type:varchar;not null"`
}

func (m *Teacher) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Teacher) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
