package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RamVellanki/habittracker/app/services"
	"github.com/gin-gonic/gin"
)

// func PostHabits(c *gin.Context) {
// 	log.Print(c.Request.Body)
// 	c.IndentedJSON(http.StatusOK, string("K"))
// }

// GetHabits godoc
// @Summary Get all active/all habits
// @Description get's active or all habits based on query
// @Tags root
// @Param status query string default "set to all to get all habits (active/inactive)"
// @Produce json
// @Success 200
// @Router /habits [get]
func GetHabits(c *gin.Context) {
	status := c.DefaultQuery("status", "")
	log.Print("status: " + status)
	getAll := false
	if status == "all" {
		getAll = true
	}
	data, err := json.Marshal(*services.GetHabits(getAll))
	if err != nil {
		fmt.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, string(data))
}

// GetHabitsById godoc
// @Summary Get all habits by Id
// @Description get's all habits based on input Id
// @Tags root
// @Param id path string true "use the id of the habit that you wish to get"
// @Produce json
// @Success 200
// @Router /habits/{id} [get]
func GetHabitsById(c *gin.Context) {
	log.Print(c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, "error: use valid integer id")
		return
	}

	log.Printf("ID Requested: %d", id)
	data, err := services.GetHabitsById(id)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	datajson, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, string(datajson))
}
