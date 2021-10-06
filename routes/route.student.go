package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mahasiswa/handlers"
	"mahasiswa/repositorys"
	"mahasiswa/services"
)

func InitializeRouteStudent(db *gorm.DB, router *gin.Engine) {

	/**
	@description All Handler Student
	*/
	repositoryStudent := repositorys.NewRepositoryStudent(db)
	serviceStudent := services.NewServiceStudent(repositoryStudent)
	handlerStudent := handlers.NewHandlerStudent(serviceStudent)

	group := router.Group("/api/v1")

	/**
	@description All Route Student
	*/

	group.POST("/student", handlerStudent.CreateHandlerStudent)
	group.GET("/student", handlerStudent.ResultsHadlerStudent)
	group.GET("/student/:id", handlerStudent.ResultHandlerStudent)
	group.DELETE("/student/:id", handlerStudent.DeleteHandlerStudent)
	group.PUT("/student/:id", handlerStudent.UpdateHandlerStudent)
}
