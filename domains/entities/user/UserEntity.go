package user

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	Id       uuid.UUID `db:"id" json:"id"`
	Email    string    `db:"email" json:"email"`
	Name     string    `db:"name" json:"name"`
	Password string    `db:"password" json:"password"`
}

func (u UserEntity) GetId() uuid.UUID {
	return u.Id
}

func (u UserEntity) GetEmail() string {
	return u.Email
}

func (u UserEntity) GetName() string {
	return u.Name
}

func (u UserEntity) GetPassword() string {
	return u.Password
}

func (u *UserEntity) SetId(id uuid.UUID) User {
	u.Id = id
	return u
}

func (u *UserEntity) SetEmail(email string) User {
	u.Email = email
	return u
}

func (u *UserEntity) SetName(name string) User {
	u.Name = name
	return u
}

func (u *UserEntity) SetPassword(password string) User {
	u.Password = password
	return u
}
