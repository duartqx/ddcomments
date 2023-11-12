package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	u "github.com/duartqx/ddcomments/domains/entities/user"
	m "github.com/duartqx/ddcomments/domains/models"
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
		Children: &[]m.Comment{},
		Creator:  &u.UserDTO{},
	}
}

func (cr CommentRepository) Create(comment m.Comment) error {
	var (
		id        uuid.UUID
		createdAt time.Time
		query     string
		values    []any
	)

	if comment.GetParentId() == uuid.Nil {
		query = `
		INSERT INTO comments (creator_id, thread_id, comment_text)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`
		values = []any{
			comment.GetCreatorId(),
			comment.GetThreadId(),
			comment.GetText(),
		}
	} else {
		query = `
			INSERT INTO comments (parent_id, creator_id, thread_id, comment_text)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at;
		`
		values = []any{
			comment.GetParentId(),
			comment.GetCreatorId(),
			comment.GetThreadId(),
			comment.GetText(),
		}
	}

	err := cr.db.QueryRow(query, values...).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	comment.SetId(id).SetCreatedAt(createdAt)

	return nil
}

func (cr CommentRepository) FindOneById(id uuid.UUID) (m.Comment, error) {
	query := `
		SELECT id, parent_id, creator_id, thread_id, comment_text, created_at
		FROM comments
		WHERE id = $1
		LIMIT 1
	`

	comment := cr.GetModel()

	err := cr.db.QueryRow(query).Scan(&comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr CommentRepository) FindAllByThreadId(id uuid.UUID) (*[]m.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByParentId(id uuid.UUID) (*[]m.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByCreatorId(id uuid.UUID) (*[]m.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) ExistsById(id uuid.UUID) (exists *bool) {
	cr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM comments WHERE id = $1);",
		id,
	).Scan(&exists)

	return exists
}
