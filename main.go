package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/middleware"
	"github.com/yoyo1025/persona-api/util"
)

var (
	db           *sql.DB
	openaiClient *openai.Client
)

func initDB() {
	var err error
	connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
	)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal("データベースへの接続に失敗しました：", err)
	}

	err = db.Ping()
	if err != nil {
			log.Fatal("データベースへの接続確認に失敗しました:", err)
	}
}

func initOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenAI APIキーが設定されていません")
	}
	openaiClient = openai.NewClient(apiKey)
}

func main() {
	fmt.Println("now server started...")
	initDB()
	database.SetDB(db)
	defer db.Close()

	initOpenAI()

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
