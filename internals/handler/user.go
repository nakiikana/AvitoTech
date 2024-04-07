package handler

import (
	"net/http"
)

func Get_Banner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ку-ку"))
}
