package model

import (
	"time"

	"github.com/google/uuid"
)

type Estate struct {
	ID        uuid.UUID
	Length    int
	Width     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tree struct {
	ID        string `json:"id"`
	EstateID  string `json:"estate_id"`
	Height    int    `json:"height"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
