package services

import (
	"fmt"

	c "github.com/duartqx/ddcomments/domains/entities/comment"
	r "github.com/duartqx/ddcomments/domains/repositories"
	"github.com/google/uuid"
)

type CommentService struct {
	commentRepository r.ICommentRepository
	threadRepository  r.IThreadRepository
}

func GetNewCommentService(
	commentRepository r.ICommentRepository,
	threadRepository r.IThreadRepository,
) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		threadRepository:  threadRepository,
	}
}

func (cs CommentService) Create(comment c.Comment) error {
	threadExists := cs.threadRepository.ExistsById(comment.GetThreadId())
	if comment.GetThreadId() == uuid.Nil || !*threadExists {
		return fmt.Errorf("Thread Not Found")
	}
	if comment.GetParentId() != uuid.Nil && !*cs.commentRepository.ExistsById(comment.GetParentId()) {
		return fmt.Errorf("Parent Not Found")
	}
	return cs.commentRepository.Create(comment)
}
