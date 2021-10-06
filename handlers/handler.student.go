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

	res, errorCode := h.service.CreateServiceStudent(&input)

	switch errorCode {
	case 409:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Npm already taken", "data": nil})
		return
	case 403:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Create new student failed", "data": nil})
		return
	default:
		ctx.JSON(http.StatusOK, gin.H{"message": "Create new student success", "data": res})
	}
}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (h *handlerStudent) ResultsHadlerStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "hello wordl"})
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
