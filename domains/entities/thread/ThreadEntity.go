package thread

import (
	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
)

type ThreadEntity struct {
	Id       uuid.UUID    `db:"id" json:"id"`
	Slug     string       `db:"slug" json:"slug"`
	Comments *[]c.Comment `json:"comments"`
}

func (t ThreadEntity) GetId() uuid.UUID {
	return t.Id
}
