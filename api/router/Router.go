package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	r "github.com/duartqx/ddcomments/infrastructure/repositories/postgres"

	lm "github.com/duartqx/ddcomments/application/middleware/logger"
	rm "github.com/duartqx/ddcomments/application/middleware/recovery"

	c "github.com/duartqx/ddcomments/api/controllers"
	s "github.com/duartqx/ddcomments/application/services"
)

type router struct {
	db     *sqlx.DB
	secret *[]byte
}

func NewRouterBuilder() *router {
	return &router{}
}

func (ro *router) SetDb(db *sqlx.DB) *router {
	ro.db = db
	return ro
}

func (ro *router) SetSecret(secret []byte) *router {
	ro.secret = &secret
	return ro
}

func (ro router) Build() *chi.Mux {

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

	router := chi.NewRouter()

	router.Use(rm.RecoveryMiddleware, lm.LoggerMiddleware)

	userSubrouter := chi.NewRouter()
	userSubrouter.Handle("/", userController)

	router.Mount("/user", userSubrouter)

	threadSubrouter := chi.NewRouter()
	threadSubrouter.Handle("/{thread_id}", threadController)

	commentSubrouter := chi.NewRouter()
	commentSubrouter.Handle("/", commentController)

	threadSubrouter.Mount("/{thread_id}/comment", commentSubrouter)

	router.Mount("/thread", threadSubrouter)

	return router
}
