package dal

import (
	"errors"

	"github.com/RamVellanki/habittracker/app/database"
	"github.com/RamVellanki/habittracker/app/models"
)

// here goes access to DB and queries
func GetHabits(getAll bool) *[]models.Habit {
	if getAll {
		return &database.Habits
	} else {
		var data []models.Habit
		for _, v := range database.Habits {
			if v.Status {
				data = append(data, v)
			}
		}
		return &data
	}
}

func GetHabitsById(id int) (*models.Habit, error) {
	for _, v := range database.Habits {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("item not found")
}
