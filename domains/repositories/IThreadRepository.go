package repositories

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type IThreadRepository interface {
	FindOneById(id uuid.UUID) (m.Thread, error)
	ExistsById(id uuid.UUID) *bool
	FindAllCommentsByThreadId(id uuid.UUID) (*[]m.Comment, error)
	Create(thread m.Thread) error
}
