package handlers

import (
	"github.com/gin-gonic/gin"

	controllers "mahasiswa/controllers/student"
)

/**
* =======================================
* 	Struct Initialize Student Database
* =======================================
 */

type handlerStudent struct {
	service controllers.Service
}

/**
* ============================
* 	Mandatory Student Method
* ============================
 */

func NewHandlerStudent(service controllers.Service) *handlerStudent {
	return &handlerStudent{service: service}
}

/**
* ========================
* 	Create Student Data
* ========================
 */

func (h *handlerStudent) CreateHandlerStudent(ctx *gin.Context) {}

/**
* ========================
* 	Results Student Data
* ========================
 */

func (h *handlerStudent) ResultsHadlerStudent(ctx *gin.Context) {}

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
