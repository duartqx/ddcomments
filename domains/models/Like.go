package models

import (
	"time"

	"github.com/google/uuid"
)

type Like interface {
	GetId() uuid.UUID
	GetUserId() uuid.UUID
	GetCommentId() uuid.UUID
	GetCreatedAt() time.Time

	GetUser() User
	GetComment() Comment

	SetId(id uuid.UUID) Like
	SetUserId(userId uuid.UUID) Like
	SetCommentId(commentId uuid.UUID) Like
	SetCreatedAt(t time.Time) Like

	SetUser(user User) Like
	SetComment(comment Comment) Like
}
