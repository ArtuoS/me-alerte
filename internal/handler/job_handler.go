package handler

import (
	"html/template"
	"net/http"

	"github.com/ArtuoS/me-alerte/internal/service"
	"github.com/ArtuoS/me-alerte/model"
	"github.com/ArtuoS/me-alerte/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobsHandler struct {
	jobService *service.JobService
}

func NewJobHandler(jobService *service.JobService) *JobsHandler {
	return &JobsHandler{
		jobService: jobService,
	}
}

func (j *JobsHandler) Get(w http.ResponseWriter, r *http.Request) {
	jobs, err := j.jobService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jobsModel := map[string][]model.Job{
		"jobs": jobs,
	}

	tmpl := template.Must(template.ParseFiles("templates/job/list_jobs.html"))
	tmpl.Execute(w, jobsModel)
}

func (j *JobsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ExtractIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	job, err := j.jobService.GetByID(objectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/job/unique_job.html"))
	tmpl.Execute(w, job)
}
