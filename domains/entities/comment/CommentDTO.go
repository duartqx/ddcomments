package comment

import (
	"time"

	"github.com/google/uuid"

	u "github.com/duartqx/ddcomments/domains/entities/user"
	m "github.com/duartqx/ddcomments/domains/models"
)

type CommentDTO struct {
	Id        uuid.UUID `db:"id" json:"id"`
	ParentId  uuid.UUID `db:"parent_id" json:"parent_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ThreadId  uuid.UUID `db:"thread_id" json:"thread_id"`
	Text      string    `db:"comment_text" json:"text"`

	Creator m.User `json:"creator"`
}

func (c CommentDTO) GetId() uuid.UUID {
	return c.Id
}

func (c CommentDTO) GetParentId() uuid.UUID {
	return c.ParentId
}

func (c CommentDTO) GetCreatorId() uuid.UUID {
	if c.Creator != nil {
		return c.Creator.GetId()
	}
	return uuid.Nil
}

func (c CommentDTO) GetThreadId() uuid.UUID {
	return c.ThreadId
}

func (c CommentDTO) GetText() string {
	return c.Text
}

func (c CommentDTO) GetCreator() m.User {
	return nil
}

func (c CommentDTO) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c CommentDTO) GetChilden() *[]m.Comment {
	return nil
}

func (c *CommentDTO) SetId(id uuid.UUID) m.Comment {
	c.Id = id
	return c
}
func (c *CommentDTO) SetParentId(parentId uuid.UUID) m.Comment {
	c.ParentId = parentId
	return c
}

func (c *CommentDTO) SetThreadId(threadId uuid.UUID) m.Comment {
	c.ThreadId = threadId
	return c
}

func (c *CommentDTO) SetCreatorId(creatorId uuid.UUID) m.Comment {
	c.Creator = &u.UserDTO{Id: creatorId}
	return c
}

func (c *CommentDTO) SetCreatedAt(createdAt time.Time) m.Comment {
	c.CreatedAt = createdAt
	return c
}

func (c *CommentDTO) SetText(text string) m.Comment {
	c.Text = text
	return c
}

func (c *CommentDTO) SetCreator(creator m.User) m.Comment {
	c.Creator = creator
	return c
}

func (c *CommentDTO) AddChildren(child ...m.Comment) m.Comment {
	return c
}
