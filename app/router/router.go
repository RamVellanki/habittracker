package router

import (
	"github.com/RamVellanki/habittracker/app/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() {
	router := gin.Default()

	router.GET("/habits", handler.GetHabits)
	router.GET("/habits/:id", handler.GetHabitsById)
	router.Run("0.0.0.0:4040")
}
