package api

import (
	"tools/internals/handler"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user_banner", handler.Get_Banner).Methods("GET")
	return r
}
