package repositorys

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"mahasiswa/models"
	"mahasiswa/schemas"
)

/**
* =======================================
* 	Interface Initialize Student Method
* =======================================
 */

type RepositoryStudent interface {
	CreateRepositoryStudent(input *schemas.Student) (*models.Student, interface{})
	ResultsRepositoryStudent() (*[]models.Student, interface{})
	ResultRepositoryStudent(input *schemas.Student) (*models.Student, interface{})
	DeleteRepositoryStudent(input *schemas.Student) (*models.Student, interface{})
	UpdateRepositoryStudent(input *schemas.Student) (*models.Student, interface{})
}

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type repositoryStudent struct {
	db *gorm.DB
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewRepositoryStudent(db *gorm.DB) *repositoryStudent {
	return &repositoryStudent{db: db}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (r *repositoryStudent) CreateRepositoryStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student models.Student
	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid
	student.Teachers = input.Teachers

	errorCode := make(chan int, 1)

	db := r.db.Model(&student)

	checkStudentNpm := db.Debug().First(&student, "npm = ?", student.Npm)

	if checkStudentNpm.RowsAffected > 0 {
		defer logrus.Error(checkStudentNpm.Error)
		errorCode <- http.StatusConflict
		return &student, <-errorCode
	}

	saveStudent := db.Debug().Create(&student).Commit()

	if saveStudent.RowsAffected < 1 {
		defer logrus.Error(saveStudent.Error)
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

func (r *repositoryStudent) ResultsRepositoryStudent() (*[]models.Student, interface{}) {
	var students []models.Student

	errorCode := make(chan int, 1)

	db := r.db.Model(&students)

	checkStudent := db.Debug().Find(&students)

	if checkStudent.RowsAffected < 1 {
		defer logrus.Error(checkStudent.Error)
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

func (r *repositoryStudent) ResultRepositoryStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student models.Student
	student.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&student)

	checkStudentById := db.Debug().First(&student)

	if checkStudentById.RowsAffected < 1 {
		defer logrus.Error(checkStudentById.Error)
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

func (r *repositoryStudent) DeleteRepositoryStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student models.Student
	student.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&student)

	checkStudentById := db.Debug().First(&student)

	if checkStudentById.RowsAffected < 1 {
		defer logrus.Error(checkStudentById.Error)
		errorCode <- http.StatusNotFound
		return &student, <-errorCode
	}

	deleteStudentById := db.Debug().Delete(&student)

	if deleteStudentById.RowsAffected < 1 {
		defer logrus.Error(deleteStudentById.Error)
		errorCode <- http.StatusForbidden
		return &student, <-errorCode
	}

	return &student, nil
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (r *repositoryStudent) UpdateRepositoryStudent(input *schemas.Student) (*models.Student, interface{}) {
	var student models.Student
	student.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&student)

	checkStudentById := db.Debug().First(&student)

	if checkStudentById.RowsAffected < 1 {
		defer logrus.Error(checkStudentById.Error)
		errorCode <- http.StatusNotFound
		return &student, <-errorCode
	}

	student.Name = input.Name
	student.Npm = input.Npm
	student.Fak = input.Fak
	student.Bid = input.Bid
	student.Teachers = input.Teachers

	updateStudent := db.Debug().Updates(&student)

	if updateStudent.RowsAffected < 1 {
		defer logrus.Error(updateStudent.Error)
		errorCode <- http.StatusForbidden
		return &student, <-errorCode
	}

	return &student, nil
}
