package main

import (
	"net/http"

	"github.com/ArtuoS/me-alerte/cmd/web/controller"
	"github.com/ArtuoS/me-alerte/internal/database"
	"github.com/ArtuoS/me-alerte/internal/repository"
	"github.com/ArtuoS/me-alerte/internal/service"
)

func main() {
	connectionString := "mongodb://localhost:27017"
	dbInstance, err := database.NewDB(connectionString)
	if err != nil {
		return
	}
	defer dbInstance.Disconnect()

	jobRepository := repository.NewJobRepository(dbInstance)
	jobService := service.NewJobService(jobRepository)
	jobController := controller.NewJobController(jobService)

	scrapDetailRepository := repository.NewScrapDetailRepository(dbInstance)
	scrapDetailService := service.NewScrapDetailService(scrapDetailRepository)
	scrapDetailController := controller.NewScrapDetailController(scrapDetailService)

	http.HandleFunc("/jobs", jobController.Get)
	http.HandleFunc("/jobs/{id}", jobController.GetByID)
	http.HandleFunc("/scrap-details", scrapDetailController.Get)

	http.ListenAndServe("127.0.0.1:8081", nil)
}
