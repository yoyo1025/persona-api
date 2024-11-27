package model

type Persona struct {
	ID         int64  `json:"persona_id" db:"persona_id"`
	Name       string `json:"name" db:"name"`
	UserID		 int64  `json:"user_id" db:"user_id"`
	Sex        string `json:"sex" db:"sex"`
	Age        int64  `json:"age" db:"age"`
	Profession string `json:"profession" db:"profession"`
	Problems   string `json:"problems" db:"problems"`
	Behavior   string `json:"behavior" db:"behavior"`
}
