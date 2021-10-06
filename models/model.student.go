package models

import (
	"time"
)

type Student struct {
	ID        uint   `gorm:"primary_key:auto_increment:unique"`
	Name      string `gorm:"type:varchar;not null"`
	Npm       uint64 `gorm:"type:bigint;not null"`
	Fak       string `gorm:"type:varchar;not null"`
	Bid       string `gorm:"type:varchar;not null"`
	TeacherId uint
	Teacher   Teacher
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Student) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Student) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
