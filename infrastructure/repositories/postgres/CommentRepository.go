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

	comments := &[]c.Comment{}

	commentsPointerMap := map[uuid.UUID]*[]c.Comment{}

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

	rows, err := cr.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := cr.GetModel()

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

		var iComment c.Comment = comment

		if sisters, ok := commentsPointerMap[iComment.GetParentId()]; !ok {
			commentsPointerMap[iComment.GetParentId()] = &[]c.Comment{iComment}
		} else {
			*sisters = append(*sisters, iComment)
		}

		if children, ok := commentsPointerMap[iComment.GetId()]; ok {
			iComment.AddChildren(*children...)
		}

		if iComment.GetParentId() == uuid.Nil {
			*comments = append(*comments, iComment)
		}
	}

	return comments, nil
}

func (cr CommentRepository) FindAllByParentId(id uuid.UUID) (*[]c.Comment, error) {
	return nil, nil
}

func (cr CommentRepository) FindAllByCreatorId(id uuid.UUID) (*[]c.Comment, error) {
	return nil, nil
}
