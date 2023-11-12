package repositories

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type ICommentRepository interface {
	Create(comment m.Comment) error
	FindOneById(id uuid.UUID) (m.Comment, error)
	ExistsById(id uuid.UUID) *bool
	FindAllByThreadId(id uuid.UUID) (*[]m.Comment, error)
	FindAllByParentId(id uuid.UUID) (*[]m.Comment, error)
	FindAllByCreatorId(id uuid.UUID) (*[]m.Comment, error)
}
