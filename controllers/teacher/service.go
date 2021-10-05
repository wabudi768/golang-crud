package controllers

import (
	"mahasiswa/models"
)

/**
* =======================================
* 	Interface Initialize Student Method
* =======================================
 */

type Service interface {
	CreateServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	ResultsServiceTeacher() (*models.EntityTeacher, interface{})
	ResultServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	DeleteServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
	UpdateServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{})
}

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type service struct {
	repository Repository
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewServiceTeacher(repository Repository) *service {
	return &service{repository: repository}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (s *service) CreateServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	res, err := s.repository.CreateRepositoryTeacher(input)
	return res, err
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (s *service) ResultsServiceTeacher() (*models.EntityTeacher, interface{}) {
	res, err := s.repository.ResultsRepositoryTeacher()
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *service) ResultServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	res, err := s.repository.ResultRepositoryTeacher(input)
	return res, err
}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (s *service) DeleteServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	res, err := s.repository.DeleteRepositoryTeacher(input)
	return res, err
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (s *service) UpdateServiceTeacher(input *InputTeacher) (*models.EntityTeacher, interface{}) {
	res, err := s.repository.UpdateRepositoryTeacher(input)
	return res, err
}
