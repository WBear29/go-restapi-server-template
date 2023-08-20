package entity

import "github.com/google/uuid"

type Sample struct {
	ID   uuid.UUID
	Name string
	TimeStamps
}
