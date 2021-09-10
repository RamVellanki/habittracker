package main

import (
	"github.com/RamVellanki/habittracker/app/db/migrations"
	database "github.com/RamVellanki/habittracker/app/db/postgres"
	"github.com/RamVellanki/habittracker/app/router"
)

// @title Habit Tracker API
// @version 1.0
// @description This is a simple API built in Go

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	database.InitDb()
	migrations.Migrate()
	router.RegisterRouter()
}
