package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	t "github.com/duartqx/ddcomments/domains/entities/thread"
	u "github.com/duartqx/ddcomments/domains/entities/user"

	m "github.com/duartqx/ddcomments/domains/models"
)

type ThreadRepository struct {
	db *sqlx.DB
}

func GetNewThreadRepository(db *sqlx.DB) *ThreadRepository {
	return &ThreadRepository{
		db: db,
	}
}

func (tr ThreadRepository) GetCommentModel() *c.CommentEntity {
	return &c.CommentEntity{
		Children: &[]m.Comment{},
		Creator:  &u.UserDTO{},
	}
}

func (tr ThreadRepository) GetThreadModel() *t.ThreadEntity {
	return &t.ThreadEntity{}
}

func (tr ThreadRepository) FindAllCommentsByThreadId(id uuid.UUID) (*[]m.Comment, error) {

	comments := &[]m.Comment{}

	query := `
		SELECT 
			-- thread info
			t.id AS thread__id,
			-- comment info
			c.id AS comment__id,
			c.comment_text AS comment__text,
			c.created_at AS comment__created_at,
			c.parent_id AS comment__parent_id,
			-- user info
			u.id AS creator__id,
			u.email AS creator__email,
			u.name AS creator__name

		FROM threads AS t

		INNER JOIN comments AS c
		ON c.thread_id = t.id

		INNER JOIN users AS u
		ON c.creator_id = u.id

		WHERE t.id = $1

		ORDER BY c.created_at DESC, c.parent_id ASC;
	`

	rows, err := tr.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := tr.GetCommentModel()

		var (
			// Creator info
			creator_id    uuid.UUID
			creator_email string
			creator_name  string
		)

		if err := rows.Scan(
			// Thread info
			&comment.ThreadId,
			// Comment info
			&comment.Id,
			&comment.Text,
			&comment.CreatedAt,
			&comment.ParentId,
			// Creator info
			&creator_id,
			&creator_email,
			&creator_name,
		); err != nil {
			return nil, err
		}

		comment.Creator.
			SetId(creator_id).
			SetName(creator_name).
			SetEmail(creator_email)

		var iComment m.Comment = comment

		*comments = append(*comments, iComment)
	}

	return comments, nil
}

func (tr ThreadRepository) FindOneById(id uuid.UUID) (m.Thread, error) {

	thread := tr.GetThreadModel()

	if err := tr.db.QueryRow(
		"SELECT * FROM thread WHERE id = $1 LIMIT 1", id,
	).Scan(&thread); err != nil {
		return nil, err
	}
	return thread, nil
}

func (tr ThreadRepository) ExistsById(id uuid.UUID) (exists *bool) {
	tr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM threads WHERE id = $1);",
		id,
	).Scan(&exists)

	return exists
}

func (tr ThreadRepository) Create(thread m.Thread) error {
	var id uuid.UUID

	if err := tr.db.QueryRow(`
			INSERT INTO threads (slug, creator_id) VALUES ($1, $2) RETURNING id
		`,
		thread.GetSlug(),
		thread.GetCreatorId(),
	).Scan(&id); err != nil {
		return err
	}

	thread.SetId(id)

	return nil
}
