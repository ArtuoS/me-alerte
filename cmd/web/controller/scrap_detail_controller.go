package controller

import (
	"html/template"
	"net/http"

	"github.com/ArtuoS/me-alerte/internal/model"
	"github.com/ArtuoS/me-alerte/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScrapDetailController struct {
	scrapDetailService *service.ScrapDetailService
}

func NewScrapDetailController(scrapDetailService *service.ScrapDetailService) *ScrapDetailController {
	return &ScrapDetailController{
		scrapDetailService: scrapDetailService,
	}
}

func FormattedID(ID primitive.ObjectID) string {
	return ID.Hex()
}

func (s *ScrapDetailController) Get(w http.ResponseWriter, r *http.Request) {
	scrapDetails, err := s.scrapDetailService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scrapDetailsModel := map[string][]model.ScrapDetail{
		"scrapDetails": scrapDetails,
	}

	funcMap := template.FuncMap{
		"formattedID": FormattedID,
	}

	tmpl := template.Must(template.ParseFiles("templates/scrap_detail/list_scrap_details.html")).Funcs(funcMap)
	tmpl.Execute(w, scrapDetailsModel)
}
