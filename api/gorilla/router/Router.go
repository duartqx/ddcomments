package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"

	lm "github.com/duartqx/ddcomments/application/middleware/logger"
	rm "github.com/duartqx/ddcomments/application/middleware/recovery"

	cc "github.com/duartqx/ddcomments/api/gorilla/controllers/comment"
	cs "github.com/duartqx/ddcomments/application/services/comment"
)

func NotFoundHandler(r *mux.Router) http.Handler {
	return r.
		NewRoute().
		BuildOnly().
		Handler(lm.LoggerMiddleware(http.HandlerFunc(http.NotFound))).
		GetHandler()
}

func MethodNotAllowedHandler(r *mux.Router) http.Handler {
	e := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

	return r.
		NewRoute().
		BuildOnly().
		Handler(lm.LoggerMiddleware(http.HandlerFunc(e))).
		GetHandler()
}

func GetMux(db *sqlx.DB) *mux.Router {

	commentRepository := r.GetNewCommentRepository(db)
	commentService := cs.GetNewCommentService(commentRepository)
	commentController := cc.GetNewCommentController(commentService)

	router := mux.NewRouter()

	router.NotFoundHandler = NotFoundHandler(router)
	router.MethodNotAllowedHandler = MethodNotAllowedHandler(router)

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	commentSubrouter := router.PathPrefix("/comment").Subrouter()

	commentSubrouter.
		Name("comment").
		Path("/{thread_id}").
		Handler(commentController).
		Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	return router
}
