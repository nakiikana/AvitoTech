package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrorBannerTagIdNotFound     = errors.New("can't get tag_id")
	ErrorBannerFeatureIdNotFound = errors.New("can't get feature_id")
	ErrorBannerIdFound           = errors.New("can't get banner id")
	ErrorReadingCreateJson       = errors.New("can't decode the json")
	ErrorReadingUpdateJson       = errors.New("can't decode the json")
	ErrorAccess                  = errors.New("no access granted")
	ErrorNoRows                  = errors.New("no rows found")
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
