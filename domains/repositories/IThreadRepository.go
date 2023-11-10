package repositories

import (
	"github.com/google/uuid"

	t "github.com/duartqx/ddcomments/domains/entities/thread"
)

type IThreadRepository interface {
	FindById(id uuid.UUID) (t.Thread, error)
}
