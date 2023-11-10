package comment

import (
	"time"

	"github.com/google/uuid"

	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type Comment interface {
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	GetCreatorId() uuid.UUID
	GetCreatedAt() time.Time
	GetThreadId() uuid.UUID
	GetText() string

	GetCreator() u.User
	GetChilden() *[]Comment

	SetId(id uuid.UUID) Comment
	SetParentId(parentId uuid.UUID) Comment
	SetCreatorId(creatorId uuid.UUID) Comment
	SetCreatedAt(t time.Time) Comment
	SetText(text string) Comment

	SetCreator(creator u.User) Comment
	AddChildren(child ...Comment) Comment
}
