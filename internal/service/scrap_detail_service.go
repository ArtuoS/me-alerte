package service

import (
	"github.com/ArtuoS/me-alerte/internal/repository"
	"github.com/ArtuoS/me-alerte/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScrapDetailService struct {
	scrapDetailRepository *repository.ScrapDetailRepository
}

func NewScrapDetailService(scrapDetailRepository *repository.ScrapDetailRepository) *ScrapDetailService {
	return &ScrapDetailService{
		scrapDetailRepository: scrapDetailRepository,
	}
}

func (r *ScrapDetailService) InsertOne(scrapDetail *model.ScrapDetail) (primitive.ObjectID, error) {
	return r.scrapDetailRepository.InsertOne(scrapDetail)
}

func (r *ScrapDetailService) UpdateOne(scrapDetail *model.ScrapDetail) error {
	return r.scrapDetailRepository.UpdateOne(scrapDetail)
}

func (r *ScrapDetailService) GetAll() ([]model.ScrapDetail, error) {
	return r.scrapDetailRepository.GetAll()
}
