package user

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type UserEntity struct {
	Id       uuid.UUID `db:"id" json:"id"`
	Email    string    `db:"email" json:"email" validate:"email,required"`
	Name     string    `db:"name" json:"name" validate:"required,min=3,max=50"`
	Password string    `db:"password" json:"password" validate:"required"`
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

func (u *UserEntity) SetId(id uuid.UUID) m.User {
	u.Id = id
	return u
}

func (u *UserEntity) SetEmail(email string) m.User {
	u.Email = email
	return u
}

func (u *UserEntity) SetName(name string) m.User {
	u.Name = name
	return u
}

func (u *UserEntity) SetPassword(password string) m.User {
	u.Password = password
	return u
}
