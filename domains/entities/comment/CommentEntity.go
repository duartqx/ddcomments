package comment

import (
	"time"

	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type CommentEntity struct {
	Id        uuid.UUID `db:"id" json:"id"`
	ParentId  uuid.UUID `db:"parent_id" json:"parent_id"`
	CreatorId uuid.UUID `db:"creator_id" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ThreadId  uuid.UUID `db:"thread_id" json:"-"`
	Text      string    `db:"comment_text" json:"text"`

	Creator  m.User       `json:"creator"`
	Children *[]m.Comment `json:"children"`
}

func (c CommentEntity) GetId() uuid.UUID {
	return c.Id
}

func (c CommentEntity) GetParentId() uuid.UUID {
	return c.ParentId
}

func (c CommentEntity) GetCreatorId() uuid.UUID {
	return c.CreatorId
}

func (c CommentEntity) GetThreadId() uuid.UUID {
	return c.ThreadId
}

func (c CommentEntity) GetText() string {
	return c.Text
}

func (c CommentEntity) GetCreator() m.User {
	return c.Creator
}

func (c CommentEntity) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c CommentEntity) GetChilden() *[]m.Comment {
	return c.Children
}

func (c *CommentEntity) SetId(id uuid.UUID) m.Comment {
	c.Id = id
	return c
}
func (c *CommentEntity) SetParentId(parentId uuid.UUID) m.Comment {
	c.ParentId = parentId
	return c
}

func (c *CommentEntity) SetThreadId(threadId uuid.UUID) m.Comment {
	c.ThreadId = threadId
	return c
}

func (c *CommentEntity) SetCreatorId(creatorId uuid.UUID) m.Comment {
	c.CreatorId = creatorId
	return c
}

func (c *CommentEntity) SetCreatedAt(createdAt time.Time) m.Comment {
	c.CreatedAt = createdAt
	return c
}

func (c *CommentEntity) SetText(text string) m.Comment {
	c.Text = text
	return c
}

func (c *CommentEntity) SetCreator(creator m.User) m.Comment {
	c.CreatorId = creator.GetId()
	c.Creator = creator
	return c
}

func (c *CommentEntity) AddChildren(child ...m.Comment) m.Comment {
	*c.Children = append(*c.Children, child...)
	return c
}
