package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	student "mahasiswa/controllers/student"
	teacher "mahasiswa/controllers/teacher"
	"mahasiswa/handlers"
)

func InitializeRouteStudent(db *gorm.DB, router *gin.Engine) {

	/**
	@description All Handler Student
	*/
	repositoryStudent := student.NewRepositoryStudent(db)
	serviceStudent := student.NewServiceStudent(repositoryStudent)
	handlerStudent := handlers.NewHandlerStudent(serviceStudent)

	/**
	@description All Handler Student
	*/
	repositoryTeacher := teacher.NewRepositoryTeacher(db)
	serviceTeacher := teacher.NewServiceTeacher(repositoryTeacher)
	handlerTeacher := handlers.NewHandlerTeacher(serviceTeacher)

	group := router.Group("/api/v1")

	/**
	@description All Route Student
	*/

	group.POST("/student", handlerStudent.CreateHandlerStudent)
	group.GET("/student", handlerStudent.ResultsHadlerStudent)
	group.GET("/student/:id", handlerStudent.ResultHandlerStudent)
	group.DELETE("/student/:id", handlerStudent.DeleteHandlerStudent)
	group.PUT("/student/:id", handlerStudent.UpdateHandlerStudent)

	/**
	@description All Route Teacher
	*/

	group.POST("/student", handlerTeacher.CreateHandlerTeacher)
	group.GET("/student", handlerTeacher.ResultsHadlerTeacher)
	group.GET("/student/:id", handlerTeacher.ResultHandlerTeacher)
	group.DELETE("/student/:id", handlerTeacher.DeleteHandlerTeacher)
	group.PUT("/student/:id", handlerTeacher.UpdateHandlerTeacher)
}
