package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"

	lm "github.com/duartqx/ddcomments/application/middleware/logger"
	rm "github.com/duartqx/ddcomments/application/middleware/recovery"

	c "github.com/duartqx/ddcomments/api/gorilla/controllers"
	s "github.com/duartqx/ddcomments/application/services"
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

	threadRepository := r.GetNewThreadRepository(db)
	threadService := s.GetNewThreadService(threadRepository)
	threadController := c.GetNewThreadController(threadService)

	commentRepository := r.GetNewCommentRepository(db)
	commentService := s.GetNewCommentService(commentRepository, threadRepository)
	commentController := c.GetNewCommentController(commentService)

	router := mux.NewRouter()

	router.NotFoundHandler = NotFoundHandler(router)
	router.MethodNotAllowedHandler = MethodNotAllowedHandler(router)

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	threadSubrouter := router.PathPrefix("/thread").Subrouter()

	threadSubrouter.
		Name("thread").
		Path("/{thread_id}").
		Handler(threadController).
		Methods(http.MethodGet)

	commentSubrouter := router.PathPrefix("/comment").Subrouter()

	commentSubrouter.
		Name("comment").
		Path("/").
		Handler(commentController).
		Methods(http.MethodPost)

	return router
}
