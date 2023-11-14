package utils

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetThreadIdFromVars(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(chi.URLParam(r, "thread_id"))
}

func GetCommentIdFromVars(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(chi.URLParam(r, "comment_id"))
}

func HasCommentIdVar(r *http.Request) bool {
	return chi.URLParam(r, "comment_id") != ""
}
