package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/handler"
	"github.com/yoyo1025/persona-api/util"
)

var (
    openaiClient *openai.Client
)

func main() {
    fmt.Println("now server started...")
    e := echo.New()

    // データベースの初期化
    database.InitDB()
    defer database.GetDB().Close()

    // OpenAIクライアントの初期化
    util.InitOpenAI(&openaiClient)
    database.SetOpenAIClient(openaiClient)

    e.GET("/persona/:userID", handler.GetPersona)
    e.POST("/persona/:userID/register", handler.RegisterPersona)
    e.GET("/conversation/:personaID", handler.GetAllMessage)
    e.POST("/conversation/:personaID", handler.PostMessage)
    e.POST("/document", handler.CreateDocument)

    e.Logger.Fatal(e.Start(":3000"))
}
