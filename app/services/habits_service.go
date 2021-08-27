package services

import (
	"github.com/RamVellanki/habittracker/app/dal"
	"github.com/RamVellanki/habittracker/app/models"
)

func GetHabits(getAll bool) *[]models.Habit {
	return dal.GetHabits(getAll)
}

func GetHabitsById(id int) (*models.Habit, error) {
	data, err := dal.GetHabitsById(id)
	return data, err
}
