package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/router"
	"github.com/yoyo1025/persona-api/util"
)

var (
    openaiClient *openai.Client
)

func main() {
    fmt.Println("now server started...")

    // データベースの初期化
    database.InitDB()
    defer database.GetDB().Close()

    // OpenAIクライアントの初期化
    util.InitOpenAI(&openaiClient)
    database.SetOpenAIClient(openaiClient)

    // ルーターを取得
    handler := router.NewRouter(openaiClient)

    // サーバーを起動
    if err := http.ListenAndServe(":3000", handler); err != nil {
        log.Fatal("サーバー起動中にエラーが発生しました:", err)
    }
}
