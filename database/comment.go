package database

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yoyo1025/persona-api/model"
	"github.com/yoyo1025/persona-api/util"
)

func GetAllCommnetsByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}

	// リクエストされたパスを取得
	path := r.URL.Path

	// パスを"/"で分割
	segments := strings.Split(path, "/")

	// パスの形式をチェック
	if len(segments) >= 3 && segments[1] == "conversation" {
		personaID, err := strconv.Atoi(segments[2])
		if err != nil {
			http.Error(w, "不正なIDフォーマットです", http.StatusBadRequest)
			return
		}

		// クエリ実行
		query := `SELECT id, user_id, persona_id, comment, is_user_comment, good FROM comment WHERE persona_id = $1`
		rows, err := db.Query(query, personaID)
		if err != nil {
			http.Error(w, "コメントの取得に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// コメントを格納するスライス
		comments := []model.Comment{}

		// クエリ結果を読み込み
		for rows.Next() {
			var id, userID, personaID int64
			var comment string
			var isUserComment, good bool

			err := rows.Scan(&id, &userID, &personaID, &comment, &isUserComment, &good)
			if err != nil {
				http.Error(w, "データの読み取りに失敗しました: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// コメントデータをマップに格納
			commentData := model.Comment{
				ID:            id,
				UserID:        userID,
				PersonaID:     personaID,
				Comment:       comment,
				IsUserComment: isUserComment,
				Good:          good,
			}
			

			// スライスに追加
			comments = append(comments, commentData)
		}

		// rows.Err()での最終的なエラーチェック
		if err := rows.Err(); err != nil {
			http.Error(w, "クエリ結果の読み取り中にエラーが発生しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// JSONでレスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(comments)
		if err != nil {
			http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
		}

	} else {
		// パスが一致しない場合は404を返す
		http.NotFound(w, r)
	}
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
		return
	}

	// リクエストされたパスを取得
	path := r.URL.Path

	// パスを"/"で分割
	segments := strings.Split(path, "/")

	// パスの形式をチェック
	if len(segments) >= 3 && segments[1] == "conversation" {
		// personaIDを取得
		personaID, err := strconv.Atoi(segments[2])
		if err != nil {
			http.Error(w, "不正なIDフォーマットです", http.StatusBadRequest)
			return
		}

		// リクエストボディからコメント情報をデコード
		var commentData struct {
			Comment       string `json:"comment"`
		}

		err = json.NewDecoder(r.Body).Decode(&commentData)
		if err != nil {
			http.Error(w, "リクエストデータのデコードに失敗しました", http.StatusBadRequest)
			return
		}

		// コメントをデータベースに挿入
		query := `
			INSERT INTO comment (user_id, persona_id, comment, is_user_comment, good)
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = db.Exec(query, 1, personaID, commentData.Comment, true, false)
		if err != nil {
			http.Error(w, "データベースへのコメント挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// 指定したペルソナIDとの会話履歴を取得
		query1 := `
			SELECT id, comment FROM comment WHERE persona_id = $1
		`
		rows, err := db.Query(query1, personaID)
		if err != nil {
			http.Error(w, "コメントの取得に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// コメントを格納するスライス
		comments := []model.Comment{}

		for rows.Next() {
			var id int64
			var comment string

			err := rows.Scan(&id, &comment)
			if err != nil {
				http.Error(w, "データの読み取りに失敗しました："+err.Error(), http.StatusInternalServerError)
				return 
			}

			commentData := model.Comment{
				ID : id,
				Comment: comment,
			}

			comments = append(comments, commentData)
		}

		// rows.Err()での最終的なエラーチェック
		if err := rows.Err(); err != nil {
			http.Error(w, "クエリ結果の読み取り中にエラーが発生しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		commentText, err := util.CreateComment(comments, openaiClient)
		if err != nil {
			http.Error(w, "AI応答の生成に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// commentテーブルに挿入
		commentQuery := `
			INSERT INTO comment (user_id, persona_id, comment, is_user_comment, good)
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = db.Exec(commentQuery, 1, personaID, commentText, false, false)
		if err != nil {
			http.Error(w, "コメントの挿入に失敗しました: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// レスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(commentText)
		if err != nil {
			http.Error(w, "レスポンスのエンコードに失敗しました: "+err.Error(), http.StatusInternalServerError)
		}
	} else {
		// パスが一致しない場合は404を返す
		http.NotFound(w, r)
	}
}
