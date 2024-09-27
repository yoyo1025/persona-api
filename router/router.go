package router

import (
	"net/http"

	"github.com/sashabaranov/go-openai"
	"github.com/yoyo1025/persona-api/database"
	"github.com/yoyo1025/persona-api/handler"
	"github.com/yoyo1025/persona-api/middleware"
	"github.com/yoyo1025/persona-api/util"
)

func NewRouter(openaiClient *openai.Client) http.Handler {
    mux := http.NewServeMux()

    // ハンドラーを登録
    mux.HandleFunc("/", database.GetPersonaArchive)
    mux.HandleFunc("/register", database.RegisterPersona)
    mux.HandleFunc("/conversation/", handler.ConversationHandler)
    mux.HandleFunc("/document", func(w http.ResponseWriter, r *http.Request) {
        util.CreateDocument(w, r, openaiClient)
    })

    // ミドルウェアの適用
    return middleware.CORS(mux)
}
