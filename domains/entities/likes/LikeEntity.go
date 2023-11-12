package likes

import (
	"time"

	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type LikeEntity struct {
	Id        uuid.UUID `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"-"`
	CommentId uuid.UUID `db:"comment_id" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	User    m.User    `json:"user"`
	Comment m.Comment `json:"comment"`
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

func (l LikeEntity) GetUser() m.User {
	return l.User
}

func (l LikeEntity) GetComment() m.Comment {
	return l.Comment
}

func (l *LikeEntity) SetId(id uuid.UUID) m.Like {
	l.Id = id
	return l
}

func (l *LikeEntity) SetCreatedAt(t time.Time) m.Like {
	l.CreatedAt = t
	return l
}

func (l *LikeEntity) SetUserId(userId uuid.UUID) m.Like {
	l.UserId = userId
	return l
}

func (l *LikeEntity) SetCommentId(commentId uuid.UUID) m.Like {
	l.CommentId = commentId
	return l
}

func (l *LikeEntity) SetUser(user m.User) m.Like {
	l.UserId = user.GetId()
	l.User = user
	return l
}

func (l *LikeEntity) SetComment(comment m.Comment) m.Like {
	l.CommentId = comment.GetId()
	l.Comment = comment
	return l
}
