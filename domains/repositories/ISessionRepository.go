package repositories

import (
	"time"

	"github.com/google/uuid"

	m "github.com/duartqx/ddcomments/domains/models"
)

type ISessionRepository interface {
	Set(key string, createdAt time.Time, userId uuid.UUID) error
	Get(key string) (m.Session, error)
	Delete(key string) error
}
