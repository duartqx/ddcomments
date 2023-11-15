package thread

import (
	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type ThreadEntity struct {
	Id        uuid.UUID    `db:"id" json:"id"`
	CreatorId uuid.UUID    `db:"creator_id" json:"-"`
	Slug      string       `db:"slug" json:"slug"`
	Comments  *[]m.Comment `json:"comments"`

	Creator m.User `json:"creator"`
}

func (t ThreadEntity) GetId() uuid.UUID {
	return t.Id
}

func (t ThreadEntity) GetCreatorId() uuid.UUID {
	return t.CreatorId
}

func (t ThreadEntity) GetCreator() m.User {
	return t.Creator
}

func (t ThreadEntity) GetSlug() string {
	return t.Slug
}

func (t *ThreadEntity) SetId(id uuid.UUID) m.Thread {
	t.Id = id
	return t
}

func (t *ThreadEntity) SetCreatorId(id uuid.UUID) m.Thread {
	t.CreatorId = id
	return t
}

func (t *ThreadEntity) SetCreator(creator m.User) m.Thread {
	t.CreatorId = creator.GetId()
	t.Creator = creator
	return t
}

func (t *ThreadEntity) SetSlug(slug string) m.Thread {
	t.Slug = slug
	return t
}
