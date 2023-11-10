package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"

	cc "github.com/duartqx/ddcomments/api/gorilla/controllers/comment"
	cs "github.com/duartqx/ddcomments/application/services/comment"
)

func GetMux(db *sqlx.DB) *mux.Router {

	commentRepository := r.GetNewCommentRepository(db)
	commentService := cs.GetNewCommentService(commentRepository)
	commentController := cc.GetNewCommentController(commentService)

	router := mux.NewRouter()

	commentSubrouter := router.PathPrefix("/comment").Subrouter()

	commentSubrouter.
		Name("comment").
		Path("/{thread_id}").
		Handler(commentController).
		Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	return router
}
