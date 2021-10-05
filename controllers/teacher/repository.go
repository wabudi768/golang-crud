package controllers

import (
	"gorm.io/gorm"

	"mahasiswa/models"
)

type Repository interface {
	CreateRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	ResultsRepositoryTeacher() (*models.EntityTeacher, interface{})
	ResultRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	DeleteRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	UpdateRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryTeacher(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	var teacher models.EntityTeacher

	teacher.Name = input.Name
	teacher.ID = input.ID

	return &teacher, nil
}

func (r *repository) ResultsRepositoryTeacher() (*models.EntityTeacher, interface{}) {
	var teacher models.EntityTeacher

	return &teacher, nil
}

func (r *repository) ResultRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	var teacher models.EntityTeacher

	return &teacher, nil
}

func (r *repository) DeleteRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	var teacher models.EntityTeacher

	return &teacher, nil
}

func (r *repository) UpdateRepositoryTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	var teacher models.EntityTeacher

	return &teacher, nil
}
