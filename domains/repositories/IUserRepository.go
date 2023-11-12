package repositories

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type IUserRepository interface {
	FindByID(id uuid.UUID) (m.User, error)
	FindByEmail(email string) (m.User, error)
}
