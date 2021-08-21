package router

import (
	"github.com/RamVellanki/habittracker/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() {
	router := gin.Default()

	router.GET("/habits", handler.GetHabits)

	router.Run("localhost:8080")
}
