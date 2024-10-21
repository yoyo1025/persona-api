package database

import "github.com/jmoiron/sqlx"

// GetDB は他のパッケージからデータベース接続を取得するための関数
func GetDB() *sqlx.DB {
	return db
}
