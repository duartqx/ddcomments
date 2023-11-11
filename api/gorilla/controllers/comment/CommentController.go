package comment

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	h "github.com/duartqx/ddcomments/application/http"
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

	var response *h.HttpResponse

	switch r.Method {
	case http.MethodGet:
		response = cc.get(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if body, err := json.Marshal(response.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(response.Status)
		w.Write(body)
	}
}

func (cc CommentController) get(r *http.Request) *h.HttpResponse {

	thread_id, err := uuid.Parse(mux.Vars(r)["thread_id"])
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: "Bad Request",
		}
	}

	comments, err := cc.commentService.GetAllCommentsByThreadId(thread_id)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusInternalServerError, Body: "Internal Server Error",
		}
	}

	results := map[string]interface{}{
		"results": comments,
	}

	return &h.HttpResponse{Status: http.StatusOK, Body: results}
}
