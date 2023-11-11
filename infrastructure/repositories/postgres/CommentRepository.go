package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type CommentRepository struct {
	db *sqlx.DB
}

func GetNewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr CommentRepository) GetModel() *c.CommentEntity {
	return &c.CommentEntity{
		Children: &[]c.Comment{},
		Creator:  &u.UserDTO{},
	}
}

func (cr CommentRepository) FindOneById(id uuid.UUID) (c.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByThreadId(id uuid.UUID) (*[]c.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByParentId(id uuid.UUID) (*[]c.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByCreatorId(id uuid.UUID) (*[]c.Comment, error) {
	return nil, nil
}
