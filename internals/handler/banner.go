package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tools/internals/models"
)

func (h *Handler) FindBanner(w http.ResponseWriter, r *http.Request) {
	input := &models.BannerGetMethod{}
	var err error
	input.TagID, err = strconv.ParseUint(r.URL.Query().Get("tag_id"), 10, 64)
	if err != nil {
		log.Printf("Banner: can't get tag_id: %v", err)
		w.WriteHeader(http.StatusBadRequest) // лучше обрабатывать ошибки
		return
	}
	input.FeatureID, err = strconv.ParseUint(r.URL.Query().Get("feature_id"), 10, 64)
	if err != nil {
		log.Printf("Banner: can't get feature_id: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	input.UseLastRevision, err = strconv.ParseBool(r.URL.Query().Get("use_last_revision")) //парсить таким образом для каждого значения??
	if err != nil {
		input.UseLastRevision = false //будет ли это работать по умолчанию
	}
	//добавить токены
	//мб добавить валидацию для объектов

	if content, err := h.service.FindBanner(input); err == sql.ErrNoRows {
		log.Printf("Couldn't find the required banner: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else { //некрасиво
		if ans, err := json.Marshal(content.Content); err != nil {
			log.Printf("FindBanner: failed marshalling the string") // content is a string??
			return
		} else {
			w.Write(ans) //не так?
		}
	}

}
func (h *Handler) Banner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("что надо??????"))
}
