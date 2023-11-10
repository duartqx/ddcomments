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

	commentsPointerMap := map[uuid.UUID]*[]c.Comment{}

	comments, err := cs.commentRepository.FindAllByThreadId(threadId)

	if err != nil {
		return nil, err
	}

	for _, cmmt := range *comments {
		if sisters, ok := commentsPointerMap[cmmt.GetParentId()]; !ok {
			commentsPointerMap[cmmt.GetParentId()] = &[]c.Comment{cmmt}
		} else {
			*sisters = append(*sisters, cmmt)
		}

		if children, ok := commentsPointerMap[cmmt.GetId()]; ok {
			cmmt.AddChildren(*children...)
		}
	}

	rootComments, _ := commentsPointerMap[uuid.Nil]

	return rootComments, nil
}
