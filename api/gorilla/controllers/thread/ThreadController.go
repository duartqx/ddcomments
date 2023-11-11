package thread

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	h "github.com/duartqx/ddcomments/application/http"
	ts "github.com/duartqx/ddcomments/application/services/thread"
)

type ThreadController struct {
	threadService *ts.ThreadService
}

func GetNewThreadController(threadService *ts.ThreadService) *ThreadController {
	return &ThreadController{
		threadService: threadService,
	}
}

func (cc ThreadController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (cc ThreadController) get(r *http.Request) *h.HttpResponse {

	thread_id, err := uuid.Parse(mux.Vars(r)["thread_id"])
	if err != nil {
		return &h.HttpResponse{
			Status: http.StatusBadRequest, Body: "Bad Request",
		}
	}

	comments, err := cc.threadService.GetAllCommentsByThreadId(thread_id)
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
