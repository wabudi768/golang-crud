package controllers

import (
	"net/http"

	"gorm.io/gorm"

	"mahasiswa/models"
)

/**
* =======================================
* 	Interface Initialize Student Method
* =======================================
 */

type Repository interface {
	CreateRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{})
	ResultsRepositoryStudent() (*models.EntityStudent, interface{})
	ResultRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{})
	DeleteRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{})
	UpdateRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{})
}

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type repository struct {
	db *gorm.DB
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewRepositoryStudent(db *gorm.DB) *repository {
	return &repository{db: db}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (r *repository) CreateRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	var student models.EntityStudent

	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid

	errorCode := make(chan int, 1)

	db := r.db.Model(&student).Begin()

	checkStudentNpm := db.Debug().Select("npm").Where("npm = ?", student.Npm).Find(&student)

	if checkStudentNpm.RowsAffected > 1 {
		errorCode <- http.StatusConflict
		return &student, <-errorCode
	}

	saveStudent := db.Debug().Create(student).First(&student).Commit()

	if saveStudent.RowsAffected < 1 {
		errorCode <- http.StatusForbidden
		return &student, <-errorCode
	}

	return &student, nil
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (r *repository) ResultsRepositoryStudent() (*models.EntityStudent, interface{}) {
	var students models.EntityStudent

	errorCode := make(chan int, 1)

	db := r.db.Model(&students).Begin()

	checkStudent := db.Debug().Select("*").Take(&students)

	if checkStudent.RowsAffected > 1 {
		errorCode <- http.StatusNotFound
		return &students, <-errorCode
	}

	return &students, nil
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (r *repository) ResultRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	var student models.EntityStudent

	errorCode := make(chan int, 1)

	db := r.db.Model(&student).Begin()

	checkStudent := db.Debug().Select("*").Where("id = ?", input.ID).Take(&student)

	if checkStudent.RowsAffected < 1 {
		errorCode <- http.StatusNotFound
		return &student, <-errorCode
	}

	return &student, nil
}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (r *repository) DeleteRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	var student models.EntityStudent

	errorCode := make(chan int, 1)

	db := r.db.Model(&student).Begin()

	checkStudent := db.Debug().Select("*").Where("id = ?", input.ID).Delete(&student)

	if checkStudent.RowsAffected < 1 {
		errorCode <- http.StatusNotFound
		return &student, <-errorCode
	}

	return &student, nil
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (r *repository) UpdateRepositoryStudent(input *InputStudent) (*models.EntityStudent, interface{}) {
	var student models.EntityStudent

	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid

	errorCode := make(chan int, 1)

	db := r.db.Model(&student).Begin()

	checkStudent := db.Debug().Select("*").Where("id = ?", input.ID).Delete(&student)

	if checkStudent.RowsAffected < 1 {
		errorCode <- http.StatusNotFound
		return &student, <-errorCode
	}

	updateStudent := db.Debug().Select("*").Where("id = ?", input.ID).Updates(&student)

	if updateStudent.RowsAffected < 1 {
		errorCode <- http.StatusForbidden
		return &student, <-errorCode
	}

	return &student, nil
}
