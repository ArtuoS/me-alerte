package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScrapDetail struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Date         time.Time          `bson:"date" json:"date"`
	JobsFound    int                `bson:"jobs_found" json:"jobsFound"`
	JobsInserted int                `bson:"jobs_inserted" json:"jobsInserted"`
}
