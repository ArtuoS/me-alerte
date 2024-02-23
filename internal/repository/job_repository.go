package repository

import (
	"context"

	"github.com/ArtuoS/me-alerte/internal/database"
	"github.com/ArtuoS/me-alerte/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobRepository struct {
	Db *database.DB
}

func NewJobRepository(db *database.DB) *JobRepository {
	return &JobRepository{
		Db: db,
	}
}

func (j *JobRepository) InsertMany(jobs []interface{}) ([]interface{}, error) {
	collection := j.Db.GetDatabase().Collection("jobs")

	options := options.InsertMany()
	options.SetOrdered(false)
	options.SetBypassDocumentValidation(true)

	result, err := collection.InsertMany(context.TODO(), jobs, options)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (j *JobRepository) GetAll() ([]model.Job, error) {
	collection := j.Db.GetDatabase().Collection("jobs")
	var jobs []model.Job
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (j *JobRepository) GetByID(id primitive.ObjectID) (model.Job, error) {
	collection := j.Db.GetDatabase().Collection("jobs")
	var job model.Job
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&job)
	if err != nil {
		return model.Job{}, err
	}

	return job, nil
}
