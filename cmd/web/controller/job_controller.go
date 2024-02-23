package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/ArtuoS/me-alerte/internal/model"
	"github.com/ArtuoS/me-alerte/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobsController struct {
	jobService *service.JobService
}

func NewJobController(jobService *service.JobService) *JobsController {
	return &JobsController{
		jobService: jobService,
	}
}

func (j *JobsController) Get(w http.ResponseWriter, r *http.Request) {
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

func (j *JobsController) GetByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[2]
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

	fmt.Println(os.Getwd())
	tmpl := template.Must(template.ParseFiles("templates/job/unique_job.html"))
	tmpl.Execute(w, job)
}
