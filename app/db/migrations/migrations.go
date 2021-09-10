package migrations

import (
	database "github.com/RamVellanki/habittracker/app/db/postgres"
	"github.com/RamVellanki/habittracker/app/models"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Habit{})
}
