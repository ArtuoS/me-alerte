package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	RemoteID       int64              `bson:"remote_id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	CareerPageName string             `bson:"career_page_name" json:"careerPageName"`
	Description    string             `bson:"description" json:"description"`
	JobUrl         string             `bson:"job_url" json:"jobUrl"`
	Workplace      string             `bson:"workplace" json:"workplace"`
	City           string             `bson:"city" json:"city"`
	State          string             `bson:"state" json:"state"`
	Country        string             `bson:"country" json:"country"`
	IsRemoteWork   bool               `bson:"is_remote_work" json:"isRemoteWork"`
	PublishedDate  time.Time          `bson:"published_date" json:"publishedDate"`
	ScrapDetailID  primitive.ObjectID `bson:"scrap_detail_id" json:"scrapDetailID"`
}

func (j *Job) FormattedID() string {
	return j.ID.Hex()
}
