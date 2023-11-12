package repositories

import (
	"time"

	m "github.com/duartqx/ddcomments/domains/models"
	"github.com/google/uuid"
)

type ISessionRepository interface {
	Set(key string, createdAt time.Time, userId uuid.UUID) error
	Get(key string) (m.Session, error)
	Delete(key string) error
}
