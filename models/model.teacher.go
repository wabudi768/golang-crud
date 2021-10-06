package models

import (
	"time"
)

type Teacher struct {
	ID        uint   `gorm:"primary_key:auto_increment:unique"`
	Name      string `gorm:"type:varchar; not null"`
	Matkul    string `gorm:"type:varchar; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Teacher) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Teacher) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
