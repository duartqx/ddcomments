package likes

import (
	"time"

	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type LikeEntity struct {
	Id        uuid.UUID `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"-"`
	CommentId uuid.UUID `db:"comment_id" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	User    u.User    `json:"user"`
	Comment c.Comment `json:"comment"`
}

func (l LikeEntity) GetId() uuid.UUID {
	return l.Id
}

func (l LikeEntity) GetUserId() uuid.UUID {
	return l.UserId
}

func (l LikeEntity) GetCommentId() uuid.UUID {
	return l.CommentId
}

func (l LikeEntity) GetCreatedAt() time.Time {
	return l.CreatedAt
}

func (l LikeEntity) GetUser() u.User {
	return l.User
}

func (l LikeEntity) GetComment() c.Comment {
	return l.Comment
}

func (l *LikeEntity) SetId(id uuid.UUID) Like {
	l.Id = id
	return l
}

func (l *LikeEntity) SetCreatedAt(t time.Time) Like {
	l.CreatedAt = t
	return l
}

func (l *LikeEntity) SetUserId(userId uuid.UUID) Like {
	l.UserId = userId
	return l
}

func (l *LikeEntity) SetCommentId(commentId uuid.UUID) Like {
	l.CommentId = commentId
	return l
}

func (l *LikeEntity) SetUser(user u.User) Like {
	l.UserId = user.GetId()
	l.User = user
	return l
}

func (l *LikeEntity) SetComment(comment c.Comment) Like {
	l.CommentId = comment.GetId()
	l.Comment = comment
	return l
}
