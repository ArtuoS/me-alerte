package service

import (
	"github.com/ArtuoS/me-alerte/internal/model"
	"github.com/ArtuoS/me-alerte/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	jobRepository *repository.JobRepository
}

func NewJobService(jobRepository *repository.JobRepository) *JobService {
	return &JobService{
		jobRepository: jobRepository,
	}
}

func (j *JobService) InsertMany(jobs []interface{}) ([]interface{}, error) {
	return j.jobRepository.InsertMany(jobs)
}

func (j *JobService) GetAll() ([]model.Job, error) {
	return j.jobRepository.GetAll()
}

func (j *JobService) GetByID(id primitive.ObjectID) (model.Job, error) {
	return j.jobRepository.GetByID(id)
}
