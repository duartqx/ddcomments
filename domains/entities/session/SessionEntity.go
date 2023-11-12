package session

import (
	"time"

	"github.com/google/uuid"
)

type SessionModel struct {
	Token      string    `db:"token"`
	UserId     uuid.UUID `db:"user_id"`
	CreationAt time.Time `db:"creation_at"`
}

func (sm SessionModel) GetToken() string {
	return sm.Token
}
