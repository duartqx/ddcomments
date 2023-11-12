package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/duartqx/ddcomments/api/gorilla/utils"

	h "github.com/duartqx/ddcomments/application/http"
	s "github.com/duartqx/ddcomments/application/services"
)

type ThreadController struct {
	threadService *s.ThreadService
}

func GetNewThreadController(threadService *s.ThreadService) *ThreadController {
	return &ThreadController{
		threadService: threadService,
	}
}

func (tc ThreadController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var response *h.HttpResponse

	switch r.Method {
	case http.MethodGet:
		response = tc.get(r)
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

func (tc ThreadController) getThreadId(r *http.Request) (uuid.UUID, error) {

	threadId, err := utils.GetThreadIdFromVars(r)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Bad Request")
	}
	exists := tc.threadService.ExistsById(threadId)
	if !*exists {
		return uuid.Nil, fmt.Errorf("Not Found")
	}
	return threadId, nil
}

func (tc ThreadController) get(r *http.Request) *h.HttpResponse {

	threadId, err := tc.getThreadId(r)
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusNotFound, Body: err.Error(),
		}
	}

	comments, err := tc.threadService.GetAllCommentsByThreadId(threadId)
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
