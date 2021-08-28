package main

import (
	"github.com/RamVellanki/habittracker/app/router"
)

// @title Habit Tracker API
// @version 1.0
// @description This is a simple API built in Go

// @host localhost:4040
// @BasePath /
// @schemes http
func main() {
	router.RegisterRouter()
}
