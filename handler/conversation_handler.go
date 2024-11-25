package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yoyo1025/persona-api/repository"
	"github.com/yoyo1025/persona-api/util"
)

func GetAllMessage(c echo.Context) error {
	personaID, err := strconv.ParseInt(c.Param("personaID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}
	messages, err := repository.GetAllMessageByID(personaID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, messages)
}
     


func PostMessage(c echo.Context) error {
	personaID, err := strconv.ParseInt(c.Param("personaID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}
	userID, err := strconv.ParseInt(c.FormValue("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}
	comment := c.FormValue("comment")
	err = repository.AddMessageByID(personaID, userID, comment, true, false)
	if err != nil {
		return err
	}
	comments, err := repository.GetAllMessageByID(personaID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}
	ResponseComment, err := util.CreateComment(comments, openaiClient)
	if err != nil {
		return err
	}
	err = repository.AddMessageByID(personaID, userID, ResponseComment, false, false)
	if err != nil {
		return err
	}
	return nil
}