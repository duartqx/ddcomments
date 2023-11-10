package repositories

import (
	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
)

type ICommentRepository interface {
	FindOneById(id uuid.UUID) (c.Comment, error)
	FindAllByThreadId(id uuid.UUID) (*[]c.Comment, error)
	FindAllByParentId(id uuid.UUID) (*[]c.Comment, error)
	FindAllByCreatorId(id uuid.UUID) (*[]c.Comment, error)
}
