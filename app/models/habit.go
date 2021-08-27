package models

type Habit struct {
	ID          int    `json:"id"`
	Summary     string `json:"summary"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
	How         string `json:"how"`
	Status      bool   `json:"status"`
}
