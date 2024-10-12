package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetPersonaArchive(c echo.Context) error {
	userID := c.Param("userID")
	fmt.Printf("userID: %s\n", userID)
	return nil
}

func RegisterPersona(c echo.Context) error {
	userID := c.Param("userID")
	fmt.Printf("userID: %s\n", userID)
	return nil
}