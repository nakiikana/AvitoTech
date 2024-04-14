package handler

import (
	"net/http"
	"tools/internals/handler/middleware"
	"tools/internals/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// func (h *Handler) InitRoutes() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/user_banner", h.FindBanner).Methods("GET")
// 	r.HandleFunc("/banner", h.CreateBanner).Methods("POST")
// 	r.HandleFunc("/banner/{id}", h.DeleteBanner).Methods("DELETE")
// 	r.HandleFunc("/banner/{id}", h.UpdateBanner).Methods("PATCH")
// 	r.HandleFunc("/banner", h.GetFilteredBanner).Methods("GET")
// 	return r
// }

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/user_banner", middleware.UserAuth(http.HandlerFunc(h.FindBanner))).Methods("GET")
	r.Handle("/banner", middleware.AdminAuth(http.HandlerFunc(h.CreateBanner))).Methods("POST")
	r.Handle("/banner/{id}", middleware.AdminAuth(http.HandlerFunc(h.DeleteBanner))).Methods("DELETE")
	r.Handle("/banner/{id}", middleware.AdminAuth(http.HandlerFunc(h.UpdateBanner))).Methods("PATCH")
	r.Handle("/banner", middleware.AdminAuth(http.HandlerFunc(h.GetFilteredBanner))).Methods("GET")

	r.Use(middleware.GenToken)
	return r
}
