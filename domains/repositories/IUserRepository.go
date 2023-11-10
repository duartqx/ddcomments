package repositories

import (
	"github.com/google/uuid"

	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type IUserRepository interface {
	FindByID(id uuid.UUID) (u.User, error)
	FindByEmail(email string) (u.User, error)
}
