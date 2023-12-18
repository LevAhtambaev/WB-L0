package handlers

import "net/http"

type Handler interface {
	GetOrder(w http.ResponseWriter, r *http.Request)
}
