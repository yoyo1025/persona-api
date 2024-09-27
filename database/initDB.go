package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitDB() {
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