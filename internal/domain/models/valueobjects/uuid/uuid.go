package uuid

import "github.com/google/uuid"

type UUIDGenerator func() (uuid.UUID, error)

var DefaultGenerator UUIDGenerator = uuid.NewV7
