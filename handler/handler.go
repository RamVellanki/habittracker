package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHabits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to habit tracker!"})
}
