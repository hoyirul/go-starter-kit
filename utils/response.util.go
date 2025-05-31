package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode 	int			`json:"status_code"`
	Success 		bool		`json:"success"`
	Message 		string	`json:"message"`
	Data 				any 		`json:"data,omitempty"`
}

func RespondWithSuccess(c *gin.Context, statusCode int, message string, data any) {
	if data == nil {
		data = gin.H{}
	}

	response := APIResponse{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data:       data,
	}

	c.JSON(statusCode, response)
}

func RespondWithError(c *gin.Context, statusCode int, message string) {
	response := APIResponse{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
	}
	c.JSON(statusCode, response)
}

func RespondWithValidationErrors(c *gin.Context, statusCode int, err map[string]string) {
	response := APIResponse{
		StatusCode: 400,
		Success:    false,
		Message:    "Validation error",
		Data:       err,
	}
	c.JSON(400, response)
}
