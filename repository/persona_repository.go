package repository

import (
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/model"
)

func GetPersonaByUserID(userID int64) ([]model.Persona, error) {
	db := database.GetDB()
	query := `SELECT id AS persona_id, user_id, name, problems FROM persona WHERE user_id = $1`

  var personas []model.Persona
  err := db.Select(&personas, query, userID)
  if err != nil {
		return nil, err
  }
	return personas, nil
}

func CreatePersona(name string, userID int64, sex string, age int64, profession string, problems string, behavior string) (int64, error) {
	db := database.GetDB()
	var personaID int64
	query := `
		INSERT INTO persona (name, user_id, sex, age, profession, problems, behavior)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	err := db.QueryRow(query, 
		name, 
		userID, 
		sex, 
		age, 
		profession, 
		problems, 
		behavior,
	).Scan(&personaID)

	if err != nil {
		return 0, err
	}

	return personaID, nil
}