package services

import (
	"fmt"

	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
	r "github.com/duartqx/ddcomments/domains/repositories"
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

func (cs CommentService) Create(comment m.Comment) error {
	threadExists := cs.threadRepository.ExistsById(comment.GetThreadId())
	if comment.GetThreadId() == uuid.Nil || !*threadExists {
		return fmt.Errorf("Thread Not Found")
	}
	if comment.GetParentId() != uuid.Nil && !*cs.commentRepository.ExistsById(comment.GetParentId()) {
		return fmt.Errorf("Parent Not Found")
	}
	return cs.commentRepository.Create(comment)
}

func (cs CommentService) FindThreadFromId(id uuid.UUID) (m.Thread, error) {
	return cs.threadRepository.FindOneById(id)
}

func (cs CommentService) FindCommentFromId(id uuid.UUID) (m.Comment, error) {
	return cs.commentRepository.FindOneById(id)
}

func (cs CommentService) ThreadExistsById(id uuid.UUID) *bool {
	return cs.threadRepository.ExistsById(id)
}

func (cs CommentService) CommentExistsById(id uuid.UUID) *bool {
	return cs.commentRepository.ExistsById(id)
}
