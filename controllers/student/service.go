package controllers

import "mahasiswa/models"

/**
* =======================================
* 	Interface Initialize Student Method
* =======================================
 */

type Service interface {
	CreateServiceStudent(input *InputStudent) (*models.EntityStudent, interface{})
	ResultsServiceStudent() (*models.EntityStudent, interface{})
	ResultServiceStudent(input *InputStudent) (*models.EntityStudent, interface{})
	DeleteServiceStudent(input *InputStudent) (*models.EntityStudent, interface{})
	UpdateServiceStudent(input *InputStudent) (*models.EntityStudent, interface{})
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

func NewServiceStudent(repository Repository) *service {
	return &service{repository: repository}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (s *service) CreateServiceStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	res, err := s.repository.CreateRepositoryStudent(input)
	return res, err
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (s *service) ResultsServiceStudent() (*models.EntityStudent, interface{}) {
	res, err := s.repository.ResultsRepositoryStudent()
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *service) ResultServiceStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	res, err := s.repository.ResultRepositoryStudent(input)
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *service) DeleteServiceStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	res, err := s.repository.DeleteRepositoryStudent(input)
	return res, err
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (s *service) UpdateServiceStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	res, err := s.repository.UpdateRepositoryStudent(input)
	return res, err
}
