package repository

import (
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/model"
)

func CreateFirstMessage(userID int64, personaID int64, comment string, isUserComment bool, good bool) error {
	db := database.GetDB()
	query := `
		INSERT INTO comment (user_id, persona_id, comment, is_user_comment, good)
		VALUES ($1, $2, $3, $4, $5)
	`
	err := db.QueryRow(query, userID, personaID, comment, isUserComment, good)
	if err != nil {
		return err.Err()
	}
	return nil
}

func GetAllMessageByID(personaID int64)([]model.Comment, error) {
	db := database.GetDB()
	query := `SELECT 
							id, 
							user_id, 
							persona_id, 
							comment, 
							is_user_comment, 
							good
						FROM comment
						WHERE persona_id = $1`
	var comments []model.Comment
	err := db.Select(&comments, query, personaID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func AddMessageByID(personaID int64, userID int64, comment string, isUserComment bool, good bool) error {
	db := database.GetDB()
	query := `
		INSERT INTO comment (user_id, persona_id, comment, is_user_comment, good)
		VALUES ($1, $2, $3, $4, $5)
	`
	err := db.QueryRow(query, userID, personaID, comment, isUserComment, good)
	if err != nil {
		return err.Err()
	}
	return nil
}

