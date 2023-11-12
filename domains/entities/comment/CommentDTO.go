package comment

import (
	"time"

	"github.com/google/uuid"

	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type CommentDTO struct {
	Id        uuid.UUID `db:"id" json:"id"`
	ParentId  uuid.UUID `db:"parent_id" json:"parent_id"`
	CreatorId uuid.UUID `db:"creator_id" json:"creator_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ThreadId  uuid.UUID `db:"thread_id" json:"thread_id"`
	Text      string    `db:"comment_text" json:"text"`
}

func (c CommentDTO) GetId() uuid.UUID {
	return c.Id
}

func (c CommentDTO) GetParentId() uuid.UUID {
	return c.ParentId
}

func (c CommentDTO) GetCreatorId() uuid.UUID {
	return c.CreatorId
}

func (c CommentDTO) GetThreadId() uuid.UUID {
	return c.ThreadId
}

func (c CommentDTO) GetText() string {
	return c.Text
}

func (c CommentDTO) GetCreator() u.User {
	return nil
}

func (c CommentDTO) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c CommentDTO) GetChilden() *[]Comment {
	return nil
}

func (c *CommentDTO) SetId(id uuid.UUID) Comment {
	c.Id = id
	return c
}
func (c *CommentDTO) SetParentId(parentId uuid.UUID) Comment {
	c.ParentId = parentId
	return c
}

func (c *CommentDTO) SetCreatorId(creatorId uuid.UUID) Comment {
	c.CreatorId = creatorId
	return c
}

func (c *CommentDTO) SetCreatedAt(createdAt time.Time) Comment {
	c.CreatedAt = createdAt
	return c
}

func (c *CommentDTO) SetText(text string) Comment {
	c.Text = text
	return c
}

func (c *CommentDTO) SetCreator(creator u.User) Comment {
	return c
}

func (c *CommentDTO) AddChildren(child ...Comment) Comment {
	return c
}
