package thread

import "github.com/google/uuid"

type Thread interface {
	GetId() uuid.UUID
}
