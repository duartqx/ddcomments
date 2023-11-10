package user

import (
	"github.com/google/uuid"
)

type UserDTO struct {
	Id    uuid.UUID `db:"id" json:"id"`
	Email string    `db:"email" json:"email"`
	Name  string    `db:"name" json:"name"`
}

func (u UserDTO) GetId() uuid.UUID {
	return u.Id
}

func (u UserDTO) GetEmail() string {
	return u.Email
}

func (u UserDTO) GetName() string {
	return u.Name
}

func (u UserDTO) GetPassword() string {
	return ""
}

func (u *UserDTO) SetId(id uuid.UUID) User {
	u.Id = id
	return u
}

func (u *UserDTO) SetEmail(email string) User {
	u.Email = email
	return u
}

func (u *UserDTO) SetName(name string) User {
	u.Name = name
	return u
}

func (u *UserDTO) SetPassword(password string) User {
	return u
}
