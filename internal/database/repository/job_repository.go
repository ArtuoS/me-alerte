package repository

import (
	"context"

	"github.com/ArtuoS/me-alerte/internal/database"
	"github.com/ArtuoS/me-alerte/internal/model"
)

type JobRepository struct {
	Db *database.DB
}

func Create(db *database.DB) *JobRepository {
	return &JobRepository{
		Db: db,
	}
}

func (j *JobRepository) Add(job *model.Job) error {
	collection := j.Db.GetDatabase().Collection("jobs")
	_, err := collection.InsertOne(context.TODO(), job)
	if err != nil {
		return err
	}
	return nil
}
