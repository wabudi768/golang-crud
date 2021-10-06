package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mahasiswa/helpers"
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
		helpers.APIErrorResponse(ctx, http.StatusConflict, "Npm already taken", res)
		return
	case 403:
		helpers.APIErrorResponse(ctx, http.StatusForbidden, "Create new student failed", res)
		return
	default:
		helpers.APISuccessResponse(ctx, http.StatusCreated, "Create new student success", nil)
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
		helpers.APIErrorResponse(ctx, http.StatusNotFound, "Student data not found", res)
		return
	default:
		helpers.APISuccessResponse(ctx, http.StatusOK, "Already OK", res)
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
