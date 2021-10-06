package models

import (
	"time"
)

type Teacher struct {
	ID        uint      `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar;not null"`
	Matkul    string    `json:"matkul" gorm:"type:varchar;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestampz"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestampz"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestampz"`
}

func (m *Teacher) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Teacher) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
