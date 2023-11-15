package services

import (
	"fmt"

	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
	r "github.com/duartqx/ddcomments/domains/repositories"
)

type ThreadService struct {
	threadRepository r.IThreadRepository
}

func GetNewThreadService(threadRepository r.IThreadRepository) *ThreadService {
	return &ThreadService{
		threadRepository: threadRepository,
	}
}

func (cs ThreadService) GetAllCommentsByThreadId(threadId uuid.UUID) (*[]m.Comment, error) {

	commentsPointerMap := map[uuid.UUID]*[]m.Comment{}

	comments, err := cs.threadRepository.FindAllCommentsByThreadId(threadId)

	if err != nil {
		return nil, err
	}

	for _, cmmt := range *comments {
		if sisters, ok := commentsPointerMap[cmmt.GetParentId()]; !ok {
			commentsPointerMap[cmmt.GetParentId()] = &[]m.Comment{cmmt}
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

func (cs ThreadService) ExistsById(threadId uuid.UUID) *bool {
	return cs.threadRepository.ExistsById(threadId)
}

func (cs ThreadService) Create(thread m.Thread) error {
	if thread.GetSlug() == "" || thread.GetCreatorId() == uuid.Nil {
		return fmt.Errorf("Invalid Thread")
	}
	return cs.threadRepository.Create(thread)
}
