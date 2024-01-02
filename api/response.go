package api

import "github.com/gin-gonic/gin"

func errorFormat(message string) gin.H {
	return gin.H{
		"type": "error",
		"message": message,
	}
}