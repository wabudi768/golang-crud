package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string      `gorm:"type:varchar;not null"`
	Npm      uint64      `gorm:"type:bigint;not null"`
	Fak      string      `gorm:"type:varchar;not null"`
	Bid      string      `gorm:"type:varchar;not null"`
	Teachers interface{} `gorm:"type:json"`
}

func (m *Student) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Student) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
