package handler

import (
	"tools/internals/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user_banner", h.Get_Banner).Methods("GET")
	r.HandleFunc("/banner", h.Banner).Methods("GET")
	return r
}
