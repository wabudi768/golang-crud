package models

import (
	"time"
)

type EntityTeacher struct {
	ID        int             `gorm:"primary_key:auto_increment:unique"`
	Name      string          `gorm:"type:varchar; not null"`
	Matkul    string          `gorm:"type:varchar; not null"`
	Students  []EntityStudent `gorm:"foreignkey:TeacherID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *EntityTeacher) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *EntityTeacher) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
