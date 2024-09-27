package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/middleware"
	"github.com/yoyo1025/persona-api/util"
)

var (
	openaiClient *openai.Client
)

func main() {
	fmt.Println("now server started...")
	database.InitDB()
	defer database.GetDB().Close()

	// initOpenAI()
	util.InitOpenAI(openaiClient)

	// OpenAIクライアントをデータベースパッケージに渡す
	database.SetOpenAIClient(openaiClient)

	// マルチプレクサを作成
	mux := http.NewServeMux()

	// ハンドラーを登録
	mux.HandleFunc("/", database.GetPersonaArchive)
	mux.HandleFunc("/register", database.RegisterPersona)
	mux.HandleFunc("/conversation/", ConversationHandler)
	mux.HandleFunc("/document", func(w http.ResponseWriter, r *http.Request) {
		util.CreateDocument(w, r, openaiClient)
	})

	// CORSミドルウェアを適用
	handler := middleware.CORS(mux)
	
	// サーバーを起動
	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatal("サーバー起動中にエラーが発生しました:", err)
	}
}
