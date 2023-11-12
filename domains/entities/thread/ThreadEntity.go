package thread

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type ThreadEntity struct {
	Id       uuid.UUID    `db:"id" json:"id"`
	Slug     string       `db:"slug" json:"slug"`
	Comments *[]m.Comment `json:"comments"`
}

func (t ThreadEntity) GetId() uuid.UUID {
	return t.Id
}
