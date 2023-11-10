package comment

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	cs "github.com/duartqx/ddcomments/application/services/comment"
)

type CommentController struct {
	commentService *cs.CommentService
}

func GetNewCommentController(commentService *cs.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

func (cc CommentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cc.get(w, r)
	}
}
func (cc CommentController) get(w http.ResponseWriter, r *http.Request) {

	thread_id, err := uuid.Parse(mux.Vars(r)["thread_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comments, err := cc.commentService.GetAllCommentsByThreadId(thread_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	results := map[string]interface{}{
		"results": comments,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(results)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
