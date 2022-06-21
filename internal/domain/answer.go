package domain

type Answer struct {
	ID         int    `json:"id"`
	ExerciseId int    `json:"exercise_id"`
	QuestionID int    `json:"question_id"`
	UserID     int    `json:"user_id"`
	Answer     string `json:"answer"`
}