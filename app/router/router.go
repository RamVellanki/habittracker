package router

import (
	"github.com/RamVellanki/habittracker/app/handler"
	_ "github.com/RamVellanki/habittracker/docs/habittracker"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterRouter() {
	router := gin.Default()

	url := ginSwagger.URL("http://localhost:4040/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/habits", handler.GetHabits)
	router.GET("/habits/:id", handler.GetHabitsById)
	// router.POST("/habits", handler.PostHabits)
	router.Run("0.0.0.0:4040")
}
