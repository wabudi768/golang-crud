package handlers

import (
	"github.com/gin-gonic/gin"

	controllers "mahasiswa/controllers/teacher"
)

/**
* =======================================
* 	Struct Initialize Teacher Database
* =======================================
 */

type handlerTeacher struct {
	service controllers.Service
}

/**
* ============================
* 	Mandatory Teacher Method
* ============================
 */

func NewHandlerTeacher(service controllers.Service) *handlerTeacher {
	return &handlerTeacher{service: service}
}

/**
* ========================
* 	Create Teacher Data
* ========================
 */

func (h *handlerTeacher) CreateHandlerTeacher(ctx *gin.Context) {}

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
