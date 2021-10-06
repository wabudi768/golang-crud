package services

import (
	"mahasiswa/models"
	"mahasiswa/repositorys"
	"mahasiswa/schemas"
)

/**
* =======================================
* 	Interface Initialize Student Method
* =======================================
 */

type ServiceTeacher interface {
	CreateServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	ResultsServiceTeacher() (*models.Teacher, interface{})
	ResultServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	DeleteServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	UpdateServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
}

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type serviceTeacher struct {
	repository repositorys.Repository
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewServiceTeacher(repository repositorys.Repository) *serviceTeacher {
	return &serviceTeacher{repository: repository}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (s *serviceTeacher) CreateServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher schemas.Teacher
	teacher.Name = input.Name
	teacher.Matkul = input.Matkul

	res, err := s.repository.CreateRepositoryTeacher(&teacher)
	return res, err
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (s *serviceTeacher) ResultsServiceTeacher() (*models.Teacher, interface{}) {
	res, err := s.repository.ResultsRepositoryTeacher()
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *serviceTeacher) ResultServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher schemas.Teacher
	teacher.ID = input.ID

	res, err := s.repository.ResultRepositoryTeacher(&teacher)
	return res, err
}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (s *serviceTeacher) DeleteServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher schemas.Teacher
	teacher.ID = input.ID

	res, err := s.repository.DeleteRepositoryTeacher(&teacher)
	return res, err
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (s *serviceTeacher) UpdateServiceTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher schemas.Teacher
	teacher.ID = input.ID
	teacher.Name = input.Name
	teacher.Matkul = input.Matkul

	res, err := s.repository.UpdateRepositoryTeacher(input)
	return res, err
}
