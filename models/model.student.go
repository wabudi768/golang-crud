package models

import (
	"time"

	"github.com/lib/pq"
)

type Student struct {
	ID        uint          `json:"id" gorm:"primary_key:auto_increment"`
	Name      string        `json:"name" gorm:"type:varchar;not null"`
	Npm       uint64        `json:"npm" gorm:"type:bigint;not null"`
	Fak       string        `json:"fak" gorm:"type:varchar;not null"`
	Bid       string        `json:"bid" gorm:"type:varchar;not null"`
	Teachers  pq.Int32Array `json:"teachers" gorm:"type:text[]"`
	CreatedAt time.Time     `json:"created_at" gorm:"type:timestampz"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"type:timestampz"`
	DeletedAt time.Time     `json:"deleted_at" gorm:"type:timestampz"`
}

func (m *Student) BeforeCreate() {
	m.CreatedAt = time.Now().Local()
}

func (m *Student) BeforeUpdate() {
	m.UpdatedAt = time.Now().Local()
}
