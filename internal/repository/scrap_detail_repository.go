package repository

import (
	"context"

	"github.com/ArtuoS/me-alerte/internal/database"
	"github.com/ArtuoS/me-alerte/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScrapDetailRepository struct {
	Db *database.DB
}

func NewScrapDetailRepository(db *database.DB) *ScrapDetailRepository {
	return &ScrapDetailRepository{
		Db: db,
	}
}

func (r *ScrapDetailRepository) InsertOne(scrapDetail *model.ScrapDetail) (primitive.ObjectID, error) {
	collection := r.Db.GetDatabase().Collection("scrap_details")
	result, err := collection.InsertOne(context.TODO(), scrapDetail)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *ScrapDetailRepository) UpdateOne(scrapDetail *model.ScrapDetail) error {
	collection := r.Db.GetDatabase().Collection("scrap_details")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": scrapDetail.ID}, bson.M{"$set": scrapDetail})
	return err
}

func (r *ScrapDetailRepository) GetAll() ([]model.ScrapDetail, error) {
	collection := r.Db.GetDatabase().Collection("scrap_details")
	var scrapDetails []model.ScrapDetail
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &scrapDetails)
	if err != nil {
		return nil, err
	}

	return scrapDetails, nil
}
