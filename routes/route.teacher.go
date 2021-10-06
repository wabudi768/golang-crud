package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mahasiswa/handlers"
	"mahasiswa/repositorys"
	"mahasiswa/services"
)

func InitializeRouteTeacher(db *gorm.DB, router *gin.Engine) {

	/**
	@description All Handler Student
	*/
	repositoryTeacher := repositorys.NewRepositoryTeacher(db)
	serviceTeacher := services.NewServiceTeacher(repositoryTeacher)
	handlerTeacher := handlers.NewHandlerTeacher(serviceTeacher)

	group := router.Group("/api/v1")

	/**
	@description All Route Teacher
	*/

	group.POST("/teacher", handlerTeacher.CreateHandlerTeacher)
	group.GET("/teacher", handlerTeacher.ResultsHadlerTeacher)
	group.GET("/teacher/:id", handlerTeacher.ResultHandlerTeacher)
	group.DELETE("/teacher/:id", handlerTeacher.DeleteHandlerTeacher)
	group.PUT("/teacher/:id", handlerTeacher.UpdateHandlerTeacher)
}
