package models

import "github.com/google/uuid"

type Session struct {
	Id     string    `json:"session_id"`
	UserId uuid.UUID `json:"user_id"`
}
