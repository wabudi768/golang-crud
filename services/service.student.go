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

type ServiceStudent interface {
	CreateServiceStudent(input *schemas.Student) (*models.Student, interface{})
	ResultsServiceStudent() (*models.Student, interface{})
	ResultServiceStudent(input *schemas.Student) (*models.Student, interface{})
	DeleteServiceStudent(input *schemas.Student) (*models.Student, interface{})
	UpdateServiceStudent(input *schemas.Student) (*models.Student, interface{})
}

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type serviceStudent struct {
	repository repositorys.RepositoryStudent
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewServiceStudent(repository repositorys.RepositoryStudent) *serviceStudent {
	return &serviceStudent{repository: repository}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (s *serviceStudent) CreateServiceStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student schemas.Student
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid
	student.Teachers = input.Teachers

	res, err := s.repository.CreateRepositoryStudent(&student)
	return res, err
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (s *serviceStudent) ResultsServiceStudent() (*models.Student, interface{}) {
	res, err := s.repository.ResultsRepositoryStudent()
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *serviceStudent) ResultServiceStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student schemas.Student
	student.ID = input.ID

	res, err := s.repository.ResultRepositoryStudent(&student)
	return res, err
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (s *serviceStudent) DeleteServiceStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student schemas.Student
	student.ID = input.ID

	res, err := s.repository.DeleteRepositoryStudent(&student)
	return res, err
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (s *serviceStudent) UpdateServiceStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student schemas.Student
	student.ID = input.ID
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid
	student.Teachers = input.Teachers

	res, err := s.repository.UpdateRepositoryStudent(&student)
	return res, err
}
