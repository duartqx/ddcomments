package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment interface {
	GetId() uuid.UUID
	GetParentId() uuid.UUID
	GetCreatorId() uuid.UUID
	GetCreatedAt() time.Time
	GetThreadId() uuid.UUID
	GetText() string

	GetCreator() User
	GetChilden() *[]Comment

	SetId(id uuid.UUID) Comment
	SetParentId(parentId uuid.UUID) Comment
	SetCreatorId(creatorId uuid.UUID) Comment
	SetCreatedAt(t time.Time) Comment
	SetText(text string) Comment
	SetThreadId(id uuid.UUID) Comment

	SetCreator(creator User) Comment
	AddChildren(child ...Comment) Comment
}
