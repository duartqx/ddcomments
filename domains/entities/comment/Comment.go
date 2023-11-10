package comment

import (
	"github.com/google/uuid"

	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type Comment interface {
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	GetCreatorId() uuid.UUID
	GetThreadId() uuid.UUID
	GetText() string

	GetCreator() u.User
	GetParent() Comment
	GetChilden() *[]Comment

	SetId(id uuid.UUID) Comment
	SetParentId(parentId uuid.UUID) Comment
	SetCreatorId(creatorId uuid.UUID) Comment
	SetText(text string) Comment

	SetCreator(creator u.User) Comment
	SetParent(parent Comment) Comment
	SetChilden(child Comment) Comment
}
