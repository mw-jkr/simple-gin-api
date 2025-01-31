package common

import "github.com/gin-gonic/gin"

func BuildErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
