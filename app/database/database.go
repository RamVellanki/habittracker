// database will hold the actual JSON for now, later this will have DB connection and DB access stuff
package database

import "github.com/RamVellanki/habittracker/app/models"

var Habits = []models.Habit{
	{
		ID:          1,
		Summary:     "Do exercise for 30mins",
		Reason:      "Exercise keeps my healthy and this in turn helps me achieve my life goals with less hindrances",
		Description: "Do any type of exercise for 30mins daily, it could be Yoga, Running or Jumba",
		How:         "1.Keep myself motivated daily. 2.Remind myself daily.",
		Status:      false,
	},
	{
		ID:          2,
		Summary:     "Eat mindfully",
		Reason:      "Eating mindfully helps in eating right and reduces obesity, constipation and other health issues",
		Description: "Neither eat less or more, eat at right level every day for every meal this helps in making life better",
		How:         "1.Remind myself during dinner to eat less. 2. Eat dinner early. ",
		Status:      true,
	},
}
