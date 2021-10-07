package repositorys

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"mahasiswa/models"
	"mahasiswa/schemas"
)

/**
* =======================================
* 	Interface Initialize Teacher Method
* =======================================
 */

type Repository interface {
	CreateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	ResultsRepositoryTeacher() (*[]models.Teacher, interface{})
	ResultRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	DeleteRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	UpdateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
}

/**
* =======================================
* 	Struct Initialize Teacher Database
* =======================================
 */

type repositoryTeacher struct {
	db *gorm.DB
}

/**
* ============================
* 	Mandatory Teacher Method
* ============================
 */

func NewRepositoryTeacher(db *gorm.DB) *repositoryTeacher {
	return &repositoryTeacher{db: db}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (r *repositoryTeacher) CreateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.Name = input.Name
	teacher.Matkul = input.Matkul

	fmt.Printf("debug %v", teacher)

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherName := db.Debug().Select("name").Where("name = ?", teacher.Name).Find(&teacher)

	if checkTeacherName.RowsAffected > 0 {
		defer logrus.Error(checkTeacherName.Error)
		errorCode <- http.StatusConflict
		return &teacher, <-errorCode
	}

	saveTeacher := db.Debug().Create(&teacher).Commit().First(&teacher)
	if saveTeacher.RowsAffected < 1 {
		defer logrus.Error(saveTeacher.Error)
		errorCode <- http.StatusForbidden
		return &teacher, <-errorCode
	}

	return &teacher, nil
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (r *repositoryTeacher) ResultsRepositoryTeacher() (*[]models.Teacher, interface{}) {
	var teachers []models.Teacher

	errorCode := make(chan int, 1)

	db := r.db.Model(&teachers).Begin()

	checkTeacher := db.Debug().Select("*").Find(&teachers)

	if checkTeacher.RowsAffected > 0 {
		defer logrus.Error(checkTeacher.Error)
		errorCode <- http.StatusNotFound
		return &teachers, <-errorCode
	}

	return &teachers, nil
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (r *repositoryTeacher) ResultRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id = ?", teacher.ID).First(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	return &teacher, nil
}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (r *repositoryTeacher) DeleteRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id = ?", teacher.ID).Take(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	deleteTeacherById := db.Debug().Select("*").Where("id = ?", teacher.ID).Delete(&teacher)

	if deleteTeacherById.RowsAffected < 1 {
		defer logrus.Error(deleteTeacherById.Error)
		errorCode <- http.StatusForbidden
		return &teacher, <-errorCode
	}

	return &teacher, nil
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (r *repositoryTeacher) UpdateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID
	teacher.Name = input.Name

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id = ?", teacher.ID).Take(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	updateTeacher := db.Debug().Select("*").Where("id = ?", teacher.ID).Updates(&teacher)

	if updateTeacher.RowsAffected < 1 {
		defer logrus.Error(updateTeacher.Error)
		errorCode <- http.StatusForbidden
		return &teacher, <-errorCode
	}

	return &teacher, nil
}
