package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mahasiswa/schemas"
	"mahasiswa/services"
)

/**
* =======================================
* 	Struct Initialize Teacher Database
* =======================================
 */

type handlerTeacher struct {
	service services.ServiceTeacher
}

/**
* ============================
* 	Mandatory Teacher Method
* ============================
 */

func NewHandlerTeacher(service services.ServiceTeacher) *handlerTeacher {
	return &handlerTeacher{service: service}
}

/**
* ========================
* 	Create Teacher Data
* ========================
 */

func (h *handlerTeacher) CreateHandlerTeacher(ctx *gin.Context) {
	var input schemas.Teacher

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "parse json data failed"})
		return
	}

	res, errorCode := h.service.CreateServiceTeacher(&input)

	switch errorCode {
	case 409:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Name already taken", "data": nil})
		return
	case 403:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Create new teacher failed", "data": nil})
		return
	default:
		ctx.JSON(http.StatusOK, gin.H{"message": "Create new teacher success", "data": res})
	}
}

/**
* ========================
* 	Results Teacher Data
* ========================
 */

func (h *handlerTeacher) ResultsHadlerTeacher(ctx *gin.Context) {}

/**
* ========================
* 	Result Teacher Data
* ========================
 */

func (h *handlerTeacher) ResultHandlerTeacher(ctx *gin.Context) {}

/**
* ========================
* 	Delete Teacher Data
* ========================
 */

func (h *handlerTeacher) DeleteHandlerTeacher(ctx *gin.Context) {}

/**
* ========================
* 	Update Teacher Data
* ========================
 */

func (h *handlerTeacher) UpdateHandlerTeacher(ctx *gin.Context) {}
