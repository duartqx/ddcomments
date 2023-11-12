package controllers

import (
	"encoding/json"
	"net/http"

	h "github.com/duartqx/ddcomments/application/http"

	s "github.com/duartqx/ddcomments/application/services"
	c "github.com/duartqx/ddcomments/domains/entities/comment"
)

type CommentController struct {
	commentService *s.CommentService
}

func GetNewCommentController(commentService *s.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

func (cc CommentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response *h.HttpResponse

	switch r.Method {
	case http.MethodPost:
		response = cc.post(r)
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

func (cc CommentController) post(r *http.Request) *h.HttpResponse {

	var comment c.Comment = &c.CommentDTO{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: "Bad Request",
		}
	}

	err = cc.commentService.Create(comment)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: "Bad Request",
		}
	}

	return &h.HttpResponse{
		Status: http.StatusCreated, Body: comment,
	}
}
