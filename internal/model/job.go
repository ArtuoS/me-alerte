package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	Id          uuid.UUID
	Title       string
	Description string
	PublishDate time.Time
	Link        string
}
