package repositories

import (
	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	t "github.com/duartqx/ddcomments/domains/entities/thread"
)

type IThreadRepository interface {
	FindById(id uuid.UUID) (t.Thread, error)
	FindAllCommentsByThreadId(id uuid.UUID) (*[]c.Comment, error)
}
