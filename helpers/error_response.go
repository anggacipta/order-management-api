package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func RespondNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func RespondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
}

func RespondForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{"error": message})
}

func RespondInternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
