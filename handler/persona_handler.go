package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/model"
	"github.com/yoyo1025/persona-api/repository"
	"github.com/yoyo1025/persona-api/util"
)

var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}

func GetPersona(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}
	personas, err := repository.GetPersonaByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, personas)
}

func RegisterPersona(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}

	var newPersona model.Persona
	if err := c.Bind(&newPersona); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body"})
	}

	personaID, err := repository.CreatePersona(newPersona.Name, userID, newPersona.Sex, newPersona.Age, newPersona.Profession, newPersona.Problems, newPersona.Behavior)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	persona := model.Persona {
		ID: personaID,
		Name: newPersona.Name,
		UserID: userID,
		Sex: newPersona.Sex,
		Age: newPersona.Age,
		Profession: newPersona.Profession,
		Problems: newPersona.Problems,
		Behavior: newPersona.Behavior,
	}
	// OpenAI APIを使ってペルソナの現状を文章化するためのコメントを生成
	commentText, err := util.CreatePersonaFirstComment(persona, openaiClient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	err = repository.CreateFirstMessage(userID, personaID, commentText, false, false)
	return c.JSON(http.StatusOK, err)
}