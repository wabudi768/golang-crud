package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mahasiswa/schemas"
	"mahasiswa/services"
)

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type handlerStudent struct {
	service services.ServiceStudent
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewHandlerStudent(service services.ServiceStudent) *handlerStudent {
	return &handlerStudent{service: service}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (h *handlerStudent) CreateHandlerStudent(ctx *gin.Context) {
	var input schemas.Student

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "parse json data failed"})
		return
	}

	_, errorCode := h.service.CreateServiceStudent(&input)

	switch errorCode {
	case 409:
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Npm already taken"})
		return
	case 403:
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Create new student failed"})
		return
	default:
		ctx.JSON(http.StatusOK, gin.H{"message": "Create new student success"})
	}
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (h *handlerStudent) ResultsHadlerStudent(ctx *gin.Context) {
	res, errorCode := h.service.ResultsServiceStudent()

	switch errorCode {
	case 404:
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Student data not found"})
		return
	default:
		ctx.JSON(http.StatusOK, gin.H{"message": "Already Ok", "data": res})
	}
}

/**
* ========================
* 	Result Student Data
* ========================
 */

func (h *handlerStudent) ResultHandlerStudent(ctx *gin.Context) {}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (h *handlerStudent) DeleteHandlerStudent(ctx *gin.Context) {}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (h *handlerStudent) UpdateHandlerStudent(ctx *gin.Context) {}
