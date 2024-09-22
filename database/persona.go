package database

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yoyo1025/persona-api/model" // 新しいパッケージをインポート
	"github.com/yoyo1025/persona-api/util"

	"github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client

func SetOpenAIClient(client *openai.Client) {
	openaiClient = client
}

func RegisterPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}

	var persona model.Persona  // model.Persona型を使用

	// リクエストボディからJSONデータをデコード
	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "リクエストデータのデコードに失敗しました", http.StatusBadRequest)
		return
	}

	// データのバリデーション（必要に応じて追加）
	if persona.Name == "" || persona.Sex == "" || persona.Profession == "" {
		http.Error(w, "必要なフィールドが不足しています", http.StatusBadRequest)
		return
	}

	// データベースに挿入
	query := `
		INSERT INTO persona (name, user_id, sex, age, profession, problems, behavior)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	var insertedID int
	err = db.QueryRow(query, persona.Name, 1, persona.Sex, persona.Age, persona.Profession, persona.Problems, persona.Behavior).Scan(&insertedID)
	if err != nil {
		http.Error(w, "データベースへの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// OpenAI APIを使ってペルソナの現状を文章化するためのコメントを生成
	commentText, err := util.CreatePersonaFirstComment(persona, openaiClient)
	if err != nil {
		http.Error(w, "AI応答の生成に失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// commentテーブルに挿入
	commentQuery := `
		INSERT INTO comment (user_id, persona_id, comment, is_user_comment, good)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = db.Exec(commentQuery, 1, insertedID, commentText, false, false)
	if err != nil {
		http.Error(w, "コメントの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := map[string]interface{}{
		"message":    "データの挿入に成功しました",
		"id":         insertedID,
		"name":       persona.Name,
		"user_id":    1,
		"age":        strconv.Itoa(persona.Age),
		"profession": persona.Profession,
		"problems":   persona.Problems,
		"behavior":   persona.Behavior,
		"ai_comment": commentText,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetPersonaArchive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}

	// リクエストされたパスを取得
	path := r.URL.Path

	// パスを"/"で分割
	segments := strings.Split(path, "/")

	// パスの形式をチェック localhost:3000/:userID
	userID, err := strconv.Atoi(segments[1])
	if err != nil {
		http.Error(w, "不正なIDフォーマットです", http.StatusBadRequest)
		return
	}

	// クエリ実行
	query := `SELECT id AS persona_id, user_id, name, problems FROM persona WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, "コメントの取得に失敗しました: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 履歴を格納するスライス
	archives := []model.Archive{}

	for rows.Next() {
		var id, userID int64
		var name, problems string

		err := rows.Scan(&id, &userID, &name, &problems)
		if err != nil {
			http.Error(w, "データの読み取りに失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		archiveData := model.Archive{
			ID:       id,
			UserID:   userID,
			Name:     name,
			Problems: problems,
		}

		archives = append(archives, archiveData)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "クエリ結果の読み取り中にエラーが発生しました: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// JSONでレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(archives)
	if err != nil {
		http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
	}
}
