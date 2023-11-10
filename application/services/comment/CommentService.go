package comment

import (
	"github.com/google/uuid"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	r "github.com/duartqx/ddcomments/domains/repositories"
)

type CommentService struct {
	commentRepository r.ICommentRepository
}

func GetNewCommentService(commentRepository r.ICommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (cs CommentService) GetAllCommentsByThreadId(threadId uuid.UUID) (*[]c.Comment, error) {
	return cs.commentRepository.FindAllByThreadId(threadId)
}
