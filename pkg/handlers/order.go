package handlers

import (
	"LO/pkg/models"
	"LO/pkg/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

type OrderHandler struct {
	Tmpl            *template.Template
	Logger          *zap.SugaredLogger
	OrderRepository repository.CacheRepository
}

func NewOrderHandler(tmpl *template.Template, logger *zap.SugaredLogger, orderRepository repository.CacheRepository) *OrderHandler {
	return &OrderHandler{Tmpl: tmpl, Logger: logger, OrderRepository: orderRepository}
}

func (oh *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	oh.Logger.Info(r.URL.Query())
	orderUUID, err := uuid.Parse(r.URL.Query().Get("uuid"))
	if err != nil {
		e := oh.Tmpl.ExecuteTemplate(w, "error.html", struct {
			Text string
		}{
			Text: err.Error(),
		})

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		oh.Logger.Infof("Failed to parse id as uuid: %v", err)
		return
	}

	order, err := oh.OrderRepository.GetOrderByID(orderUUID)
	if err != nil {
		e := oh.Tmpl.ExecuteTemplate(w, "error.html", struct {
			Text string
		}{
			Text: err.Error(),
		})

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		oh.Logger.Infof("Failed to get order form cache: %v", err)
		return
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
