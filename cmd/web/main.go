package main

import (
	"net/http"

	"github.com/ArtuoS/me-alerte/internal/database"
	"github.com/ArtuoS/me-alerte/internal/handler"
	"github.com/ArtuoS/me-alerte/internal/repository"
	"github.com/ArtuoS/me-alerte/internal/service"
)

func main() {
	dbInstance, err := database.NewDB()
	if err != nil {
		return
	}
	defer dbInstance.Disconnect()

	jobRepository := repository.NewJobRepository(dbInstance)
	jobService := service.NewJobService(jobRepository)
	jobHandler := handler.NewJobHandler(jobService)

	scrapDetailRepository := repository.NewScrapDetailRepository(dbInstance)
	scrapDetailService := service.NewScrapDetailService(scrapDetailRepository)
	scrapDetailHandler := handler.NewScrapDetailHandler(scrapDetailService)

	http.HandleFunc("/jobs", jobHandler.Get)
	http.HandleFunc("/jobs/{id}", jobHandler.GetByID)
	http.HandleFunc("/scrap-details", scrapDetailHandler.Get)

	http.ListenAndServe("127.0.0.1:8081", nil)
}

func setupRoutes() {

}