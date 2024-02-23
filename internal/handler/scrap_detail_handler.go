package handler

import (
	"html/template"
	"net/http"

	"github.com/ArtuoS/me-alerte/internal/service"
	"github.com/ArtuoS/me-alerte/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScrapDetailHandler struct {
	scrapDetailService *service.ScrapDetailService
}

func NewScrapDetailHandler(scrapDetailService *service.ScrapDetailService) *ScrapDetailHandler {
	return &ScrapDetailHandler{
		scrapDetailService: scrapDetailService,
	}
}

func FormattedID(ID primitive.ObjectID) string {
	return ID.Hex()
}

func (s *ScrapDetailHandler) Get(w http.ResponseWriter, r *http.Request) {
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
