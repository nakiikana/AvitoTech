package handler

import (
	"net/http"
)

func (h *Handler) FindBanner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ку-ку"))

}
func (h *Handler) Banner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("что надо??????"))
}
