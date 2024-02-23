package main

import (
	"time"

	"github.com/ArtuoS/me-alerte/internal/database"
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
	scrapDetailRepository := repository.NewScrapDetailRepository(dbInstance)
	scrapDetailService := service.NewScrapDetailService(scrapDetailRepository)
	scrapService := service.NewScrapService(jobService, scrapDetailService)

	for {
		scrapService.StartScrapping()
		time.Sleep(time.Hour * 1)
	}
}
