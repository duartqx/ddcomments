package likes

import (
	"time"

	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type Like interface {
	GetId() uuid.UUID
	GetUserId() uuid.UUID
	GetCommentId() uuid.UUID
	GetCreatedAt() time.Time

	GetUser() u.User
	GetComment() c.Comment

	SetId(id uuid.UUID) Like
	SetUserId(userId uuid.UUID) Like
	SetCommentId(commentId uuid.UUID) Like
	SetCreatedAt(t time.Time) Like

	SetUser(user u.User) Like
	SetComment(comment c.Comment) Like
}
