package database

import (
	"log"

	"github.com/RamVellanki/habittracker/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDb() {
	log.Print("connecting to DB")
	var err error
	dsn := "host=localhost user=postgres password=admin123# dbname=habittracker port=54320 sslmode=disable"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Print(DBCon.Error)

	if err != nil {
		panic(err)
		panic("failed to connect to database")
	}
}

func GetHabits(getall bool) (*[]models.Habit, error) {
	var habits []models.Habit
	log.Print("status of dbcon", DBCon)
	DBCon.Where("status = ?", getall).Find(&habits)
	log.Print(habits)
	return &habits, nil
}
