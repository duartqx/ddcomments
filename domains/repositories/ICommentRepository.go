package repositories

import (
	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
)

type ICommentRepository interface {
	Create(comment c.Comment) error
	FindOneById(id uuid.UUID) (c.Comment, error)
	ExistsById(id uuid.UUID) *bool
	FindAllByThreadId(id uuid.UUID) (*[]c.Comment, error)
	FindAllByParentId(id uuid.UUID) (*[]c.Comment, error)
	FindAllByCreatorId(id uuid.UUID) (*[]c.Comment, error)
}
