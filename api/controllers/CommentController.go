package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/duartqx/ddcomments/api/utils"
	h "github.com/duartqx/ddcomments/application/http"

	s "github.com/duartqx/ddcomments/application/services"
	c "github.com/duartqx/ddcomments/domains/entities/comment"
	m "github.com/duartqx/ddcomments/domains/models"
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

// TODO inject with middleware
func (cc CommentController) getThreadId(r *http.Request) (uuid.UUID, error) {
	threadId, err := utils.GetThreadIdFromVars(r)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Bad Request")
	}
	exists := cc.commentService.ThreadExistsById(threadId)
	if !*exists {
		return uuid.Nil, fmt.Errorf("Not Found")
	}
	return threadId, nil
}

// TODO inject with middleware
func (cc CommentController) getCommentId(r *http.Request) (uuid.UUID, error) {
	commentId, err := utils.GetCommentIdFromVars(r)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Bad Request")
	}
	exists := cc.commentService.CommentExistsById(commentId)
	if !*exists {
		return uuid.Nil, fmt.Errorf("Not Found")
	}
	return commentId, nil
}

func (cc CommentController) post(r *http.Request) *h.HttpResponse {

	var comment m.Comment = &c.CommentDTO{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: "Bad Request",
		}
	}

	threadId, err := cc.getThreadId(r)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: err.Error(),
		}
	}
	comment.SetThreadId(threadId)

	if utils.HasCommentIdVar(r) {
		parentId, err := cc.getCommentId(r)
		if err != nil {
			return &h.HttpResponse{
				Status: http.StatusBadRequest, Body: err.Error(),
			}
		}
		comment.SetParentId(parentId)
	}

	// TODO inject with middleware
	cid, _ := uuid.Parse("81ad0ae3-84b5-47f9-8770-8e842ae60ce9")
	comment.SetCreatorId(cid)

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
