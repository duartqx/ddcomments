package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"

	lm "github.com/duartqx/ddcomments/application/middleware/logger"
	rm "github.com/duartqx/ddcomments/application/middleware/recovery"

	tc "github.com/duartqx/ddcomments/api/gorilla/controllers/thread"
	ts "github.com/duartqx/ddcomments/application/services/thread"
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
	threadService := ts.GetNewThreadService(threadRepository)
	threadController := tc.GetNewThreadController(threadService)

	router := mux.NewRouter()

	router.NotFoundHandler = NotFoundHandler(router)
	router.MethodNotAllowedHandler = MethodNotAllowedHandler(router)

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	commentSubrouter := router.PathPrefix("/thread").Subrouter()

	commentSubrouter.
		Name("thread").
		Path("/{thread_id}").
		Handler(threadController).
		Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	return router
}
