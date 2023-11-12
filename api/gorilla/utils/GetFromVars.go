package utils

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetThreadIdFromVars(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(mux.Vars(r)["thread_id"])
}

func GetCommentIdFromVars(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(mux.Vars(r)["comment_id"])
}
