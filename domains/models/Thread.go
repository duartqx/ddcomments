package models

import "github.com/google/uuid"

type Thread interface {
	GetId() uuid.UUID
	GetCreator() User
	GetCreatorId() uuid.UUID
	GetSlug() string

	SetId(id uuid.UUID) Thread
	SetCreator(User) Thread
	SetCreatorId(id uuid.UUID) Thread
	SetSlug(slug string) Thread
}
