package model

import (
	"time"

	"github.com/google/uuid"
)

type ScrapInfo struct {
	Id                uuid.UUID
	StartDate         time.Time
	EndDate           time.Time
	JobsFoundInSearch []uuid.UUID
}

func (s *ScrapInfo) JobsFound() int {
	return len(s.JobsFoundInSearch)
}
