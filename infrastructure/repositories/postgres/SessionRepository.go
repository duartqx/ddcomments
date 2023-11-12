package postgres

import (
	"fmt"
	"time"

	s "github.com/duartqx/ddcomments/domains/entities/session"
	m "github.com/duartqx/ddcomments/domains/models"
	"github.com/google/uuid"
)

type SessionStore map[string]m.Session

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (ss *SessionStore) Set(key string, createdAt time.Time, userId uuid.UUID) error {
	(*ss)[key] = &s.SessionModel{
		CreationAt: createdAt,
		UserId:     userId,
	}
	return nil
}

func (ss SessionStore) Get(key string) (m.Session, error) {
	session, ok := ss[key]
	if !ok {
		return nil, fmt.Errorf("Key not found")
	}
	return session, nil
}

func (ss *SessionStore) Delete(key string) error {
	delete(*ss, key)
	return nil
}
