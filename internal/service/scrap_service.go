package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/ArtuoS/me-alerte/internal/model"
	"github.com/go-resty/resty/v2"
)

const (
	baseURL      = "https://portal.api.gupy.io/api/v1/jobs"
	itemsPerPage = 10
)

type ScrapDetails struct {
	Jobs       []model.Job      `json:"data"`
	Pagination model.Pagination `json:"pagination"`
}

type ScrapService struct {
	jobService         *JobService
	scrapDetailService *ScrapDetailService
	restClient         *resty.Client
}

func NewScrapService(jobService *JobService, scrapDetailService *ScrapDetailService) *ScrapService {
	return &ScrapService{
		jobService:         jobService,
		scrapDetailService: scrapDetailService,
		restClient:         resty.New(),
	}
}

func (s *ScrapService) StartScrapping() {
	jobsChan := make(chan []model.Job)
	controlChan := make(chan bool)
	defer close(jobsChan)
	defer close(controlChan)

	go s.search(".NET", jobsChan, controlChan)

	scrapDetail := &model.ScrapDetail{
		Date:         time.Now(),
		JobsFound:    0,
		JobsInserted: 0,
	}
	id, err := s.scrapDetailService.InsertOne(scrapDetail)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case jobs := <-jobsChan:
			var jobInterfaces []interface{}
			for _, job := range jobs {
				job.ScrapDetailID = id
				jobInterfaces = append(jobInterfaces, job)
			}
			jobsInserted, _ := s.jobService.InsertMany(jobInterfaces)
			scrapDetail.JobsFound += len(jobs)
			scrapDetail.JobsInserted += len(jobsInserted)
		case <-controlChan:
			scrapDetail.ID = id
			s.scrapDetailService.UpdateOne(scrapDetail)
			log.Printf("Finished scrapping, %d jobs found", scrapDetail.JobsFound)
			return
		}
	}
}

func (s *ScrapService) search(keyword string, jobsChan chan<- []model.Job, controlChan chan<- bool) {
	limit, offset, total, runs := 10, 0, 0, 1
	for i := 0; i < runs; i++ {
		body, err := s.executeScrap(keyword, limit, offset)
		if err != nil {
			log.Fatalln("Error executing scrap:", err)
			return
		}

		var scrapInfo ScrapDetails
		if err := json.Unmarshal([]byte(body), &scrapInfo); err != nil {
			log.Fatalln("Error unmarshalling JSON:", err)
			return
		}

		jobsChan <- scrapInfo.Jobs

		if total == 0 {
			total = scrapInfo.Pagination.Total
			runs = int(math.Ceil(float64(total / itemsPerPage)))
		}

		offset += itemsPerPage
	}

	controlChan <- true
}

func (s *ScrapService) executeScrap(keyword string, limit, offset int) (string, error) {
	url := fmt.Sprintf("%s?jobName=%s&limit=%d&offset=%d", baseURL, keyword, limit, offset)
	resp, err := s.restClient.R().Get(url)

	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}

	fmt.Println("Status Code:", resp.Status())
	return string(resp.Body()), nil
}
