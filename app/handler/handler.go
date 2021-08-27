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

func GetHabitsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
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
