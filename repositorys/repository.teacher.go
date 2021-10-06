package repositorys

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"mahasiswa/models"
	"mahasiswa/schemas"
)

type Repository interface {
	CreateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	ResultsRepositoryTeacher() (*models.Teacher, interface{})
	ResultRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	DeleteRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
	UpdateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{})
}

type repositoryTeacher struct {
	db *gorm.DB
}

func NewRepositoryTeacher(db *gorm.DB) *repositoryTeacher {
	return &repositoryTeacher{db: db}
}

func (r *repositoryTeacher) CreateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.Name = input.Name
	teacher.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherName := db.Debug().Select("*").Where("name ? =", teacher.Name).Take(&teacher)

	if checkTeacherName.RowsAffected > 1 {
		defer logrus.Error(checkTeacherName.Error)
		errorCode <- http.StatusConflict
		return &teacher, errorCode
	}

	saveTeacher := db.Debug().Create(&teacher).First(&teacher).Commit()
	if saveTeacher.RowsAffected < 1 {
		defer logrus.Error(saveTeacher.Error)
		errorCode <- http.StatusForbidden
		return &teacher, errorCode
	}

	return &teacher, nil
}

func (r *repositoryTeacher) ResultsRepositoryTeacher() (*models.Teacher, interface{}) {
	var teachers models.Teacher

	errorCode := make(chan int, 1)

	db := r.db.Model(&teachers).Begin()

	checkTeacher := db.Debug().Select("*").Take(&teachers)

	if checkTeacher.RowsAffected > 1 {
		defer logrus.Error(checkTeacher.Error)
		errorCode <- http.StatusNotFound
		return &teachers, <-errorCode
	}

	return &teachers, nil
}

func (r *repositoryTeacher) ResultRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id ? =", teacher.ID).Take(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	return &teacher, nil
}

func (r *repositoryTeacher) DeleteRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id ? =", teacher.ID).Take(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	deleteTeacherById := db.Debug().Select("*").Where("id ? =", teacher.ID).Delete(&teacher)

	if deleteTeacherById.RowsAffected < 1 {
		defer logrus.Error(deleteTeacherById.Error)
		errorCode <- http.StatusForbidden
		return &teacher, <-errorCode
	}

	return &teacher, nil
}

func (r *repositoryTeacher) UpdateRepositoryTeacher(input *schemas.Teacher) (*models.Teacher, interface{}) {
	var teacher models.Teacher
	teacher.ID = input.ID
	teacher.Name = input.Name

	errorCode := make(chan int, 1)

	db := r.db.Model(&teacher).Begin()

	checkTeacherById := db.Debug().Select("*").Where("id ? =", teacher.ID).Take(&teacher)

	if checkTeacherById.RowsAffected < 1 {
		defer logrus.Error(checkTeacherById.Error)
		errorCode <- http.StatusNotFound
		return &teacher, <-errorCode
	}

	updateTeacher := db.Debug().Select("*").Where("id ? =", teacher.ID).Updates(&teacher)

	if updateTeacher.RowsAffected < 1 {
		defer logrus.Error(updateTeacher.Error)
		errorCode <- http.StatusForbidden
		return &teacher, <-errorCode
	}

	return &teacher, nil
}
