package user

import "github.com/google/uuid"

type User interface {
	GetId() uuid.UUID
	GetEmail() string
	GetName() string
	GetPassword() string

	SetId(id uuid.UUID) User
	SetEmail(email string) User
	SetName(name string) User
	SetPassword(password string) User
}
