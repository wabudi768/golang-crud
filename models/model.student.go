package models

import (
	"time"
)

type EntityStudent struct {
	ID        int    `gorm:"primary_key:auto_increment:unique"`
	Name      string `gorm:"type:varchar;not null"`
	Npm       uint64 `gorm:"type:integer;not null;unique"`
	Fak       string `gorm:"type:varchar;not null"`
	Bid       string `gorm:"type:varchar;not null"`
	TeacherID int    `gorm:"index:unique"`
	Teacher   EntityTeacher
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *EntityStudent) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *EntityStudent) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
