package handlers

import (
	"LO/pkg/models"
	"LO/pkg/repository"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	Tmpl            *template.Template
	Logger          *zap.SugaredLogger
	OrderRepository repository.Repository
}

func NewOrderHandler(template *template.Template, logger *zap.SugaredLogger, repository repository.Repository) *OrderHandler {
	return &OrderHandler{
		Tmpl:            template,
		Logger:          logger,
		OrderRepository: repository,
	}
}

func (oh *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := oh.OrderRepository.GetOrder(r.Context(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order.ID = orderID
	order.Items = []models.Items{
		{
			ID:          1,
			OrderID:     0,
			ChrtID:      0,
			TrackNumber: "3234",
			Price:       0,
			Rid:         "234",
			Name:        "234",
			Sale:        0,
			Size:        "234",
			TotalPrice:  0,
			NmID:        0,
			Brand:       "234",
			Status:      0,
		},
		{
			ID:          2,
			OrderID:     0,
			ChrtID:      0,
			TrackNumber: "3234",
			Price:       0,
			Rid:         "234",
			Name:        "234",
			Sale:        0,
			Size:        "234",
			TotalPrice:  0,
			NmID:        0,
			Brand:       "234",
			Status:      0,
		},
	}

	err = oh.Tmpl.ExecuteTemplate(w, "order.html", struct {
		Order models.OrderJSON
	}{
		Order: order,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
