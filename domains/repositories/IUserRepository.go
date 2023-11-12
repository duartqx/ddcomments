package repositories

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type IUserRepository interface {
	FindByID(id uuid.UUID) (m.User, error)
	FindByEmail(email string) (m.User, error)
	Update(user m.User) error
	All() (*[]m.User, error)
	Create(user m.User) error
	Delete(user m.User) error
	ExistsById(id uuid.UUID) *bool
	ExistsByEmail(email string) *bool
}
