package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

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

func main() {
	fmt.Println("now server started...")
	initDB()
	defer db.Close()

	// シンプルなハンドラー
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// サーバーをポート3000で起動
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("サーバー起動中にエラーが発生しました:", err)
	}

	fmt.Println("プログラムを終了します")
}
