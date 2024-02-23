package service

import (
	"github.com/ArtuoS/me-alerte/internal/repository"
	"github.com/ArtuoS/me-alerte/model"
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

func (j *JobService) GetForDisplay() ([]model.JobDisplay, error) {
	jobs, err := j.jobRepository.GetAll()
	if err != nil {
		return nil, err
	}

	jobDisplays := []model.JobDisplay{}
	for _, job := range jobs {
		jobDisplays = append(jobDisplays, model.JobDisplay{
			ID:             job.ID.Hex(),
			Name:           job.Name,
			Description:    job.Description,
			CareerPageName: job.CareerPageName,
			JobUrl:         job.JobUrl,
			IsRemoteWork:   job.IsRemoteWork,
			PublishedDate:  job.PublishedDate,
			ScrapDetailID:  job.ScrapDetailID.Hex(),
		})
	}

	return jobDisplays, nil
}
