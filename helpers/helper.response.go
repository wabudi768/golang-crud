package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponseError struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

func APISuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	successData := Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
	ctx.JSON(code, successData)
}

func APIErrorResponse(ctx *gin.Context, code int, message string, data interface{}) {
	errorData := ResponseError{
		StatusCode: code,
		Error:      message,
		Data:       data,
	}
	ctx.AbortWithStatusJSON(code, errorData)
}
