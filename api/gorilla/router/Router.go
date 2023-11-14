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

type Router struct {
	db     *sqlx.DB
	secret *[]byte
}

func NewRouterBuilder() *Router {
	return &Router{}
}

func (ro *Router) SetDb(db *sqlx.DB) *Router {
	ro.db = db
	return ro
}

func (ro *Router) SetSecret(secret []byte) *Router {
	ro.secret = &secret
	return ro
}

func (ro Router) notFoundHandler(r *mux.Router) http.Handler {
	return r.
		NewRoute().
		BuildOnly().
		Handler(lm.LoggerMiddleware(http.HandlerFunc(http.NotFound))).
		GetHandler()
}

func (ro Router) methodNotAllowedHandler(r *mux.Router) http.Handler {
	e := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

	return r.
		NewRoute().
		BuildOnly().
		Handler(lm.LoggerMiddleware(http.HandlerFunc(e))).
		GetHandler()
}

func (ro Router) Build() *mux.Router {

	userRepository := r.GetNewUserRepository(ro.db)
	userService := s.GetNewUserService(userRepository)
	userController := c.GetNewUserController(userService)

	threadRepository := r.GetNewThreadRepository(ro.db)
	threadService := s.GetNewThreadService(threadRepository)
	threadController := c.GetNewThreadController(threadService)

	commentRepository := r.GetNewCommentRepository(ro.db)
	commentService := s.GetNewCommentService(commentRepository, threadRepository)
	commentController := c.GetNewCommentController(commentService)

	// jwtAuthService := s.NewJwtAuthService(userRepository, r.NewSessionStore(), ro.secret)
	// jwtController := c.NewJwtController(jwtAuthService)

	router := mux.NewRouter()

	router.NotFoundHandler = ro.notFoundHandler(router)
	router.MethodNotAllowedHandler = ro.methodNotAllowedHandler(router)

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	userSubrouter := router.PathPrefix("/user").Subrouter()

	userSubrouter.
		Name("user").
		Path("/").
		Handler(userController).
		Methods(http.MethodPost)

	threadSubrouter := router.PathPrefix("/thread").Subrouter()

	threadSubrouter.
		Name("thread").
		Path("/{thread_id}").
		Handler(threadController).
		Methods(http.MethodGet)

	commentSubrouter := threadSubrouter.PathPrefix("/{thread_id}/comment").Subrouter()

	commentSubrouter.
		Name("comment").
		Path("/").
		Handler(commentController).
		Methods(http.MethodPost)

	return router
}
