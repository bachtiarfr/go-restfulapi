package domain

import "time"

type Exercise struct {
	ID    		int    `json:"id"`
	Title 		string `json:"name"`
	Description string `json:"description"`
	Question []Question
}
type Question struct {
	ID            int    `json:"id"`
	ExerciseID    int    `json:"exercise_id"`
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
	CreatorID     int    `json:"creator_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
