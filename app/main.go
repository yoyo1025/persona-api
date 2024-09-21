package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("now server started...")

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
