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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Parse json data failed"})
		return
	}

	res, error := h.service.CreateServiceStudent(&input)

	switch error {
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
	res, error := h.service.ResultsServiceStudent()

	switch error {
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

func (h *handlerStudent) ResultHandlerStudent(ctx *gin.Context) {

	var input schemas.Student
	input.ID = ctx.Param("id")

	res, error := h.service.ResultServiceStudent(&input)

	switch error {
	case 404:
		helpers.APIErrorResponse(ctx, http.StatusNotFound, "Student data not found", res)
		return
	default:
		helpers.APISuccessResponse(ctx, http.StatusOK, "Already OK", res)
	}
}

/**
* ========================
* 	Delete Student Data
* ========================
 */

func (h *handlerStudent) DeleteHandlerStudent(ctx *gin.Context) {
	var input schemas.Student
	input.ID = ctx.Param("id")

	res, error := h.service.DeleteServiceStudent(&input)

	switch error {
	case 404:
		helpers.APIErrorResponse(ctx, http.StatusNotFound, "StudentId is not exist", res)
		return
	case 403:
		helpers.APIErrorResponse(ctx, http.StatusForbidden, "Delete student data failed", res)
		return
	default:
		helpers.APISuccessResponse(ctx, http.StatusOK, "Delete student data success", nil)
	}
}

/**
* ========================
* 	Update Student Data
* ========================
 */

func (h *handlerStudent) UpdateHandlerStudent(ctx *gin.Context) {
	var input schemas.Student
	input.ID = ctx.Param("id")

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Parse json data failed"})
		return
	}

	res, error := h.service.UpdateServiceStudent(&input)

	switch error {
	case 404:
		helpers.APIErrorResponse(ctx, http.StatusNotFound, "StudentId is not exist", res)
		return
	case 403:
		helpers.APIErrorResponse(ctx, http.StatusForbidden, "Update new student failed", res)
		return
	default:
		helpers.APISuccessResponse(ctx, http.StatusOK, "Update new student success", nil)
	}
}
