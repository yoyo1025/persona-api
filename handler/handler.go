package handler

// import (
// 	"net/http"

// 	"github.com/yoyo1025/persona-api/database"
// )

// func ConversationHandler(w http.ResponseWriter, r *http.Request) {
//     if r.Method == http.MethodGet {
//         database.GetAllCommentsByID(w, r)
//     } else if r.Method == http.MethodPost {
//         database.PostComment(w, r)
//     } else {
//         http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
//     }
// }
