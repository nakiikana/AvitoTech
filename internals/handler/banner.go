package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tools/internals/models"

	"github.com/gorilla/mux"
)

func (h *Handler) FindBanner(w http.ResponseWriter, r *http.Request) {
	input := &models.BannerGetRequest{}
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
	//ToDo: ПРОВЕРКА ТОКЕНОВ
	//мб добавить валидацию для объектов

	if content, err := h.service.FindBanner(input); err == sql.ErrNoRows {
		log.Printf("Couldn't find the required banner: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else { //некрасиво
		// if ans, err := json.Marshal(content.Content); err != nil {
		// 	log.Printf("FindBanner: failed marshalling the string") // content is a string??
		// 	return
		// } else {
		w.Write(content.Content) //не так?
		// }
	}

}
func (h *Handler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	banner := &models.Banner{}
	err := json.NewDecoder(r.Body).Decode(banner)
	if err != nil {
		log.Printf("CreateBanner: couldn't decode the json: %v", err)
		JSONError(w, models.ErrorMessage{Error: "couldn't decode the json"}, http.StatusBadRequest)
		return
	}
	if id, err := h.service.CreateBanner(banner); err != nil {
		log.Printf("CreateBanner: couldn't insert a new banner: %v", err)
		JSONError(w, models.ErrorMessage{Error: "Couldn't insert a new banner"}, http.StatusBadRequest) // internal server err prob
		return
	} else {
		ans, err := json.Marshal(id)
		if err != nil {
			log.Printf("CreateBanner: couldn't insert a new banner: %v", err) //??
			JSONError(w, models.ErrorMessage{Error: "Internal Server Error"}, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(ans)
		//ToDo : ТОКЕНЫ - ПОЛЬЗОВАТЕЛЬ НЕ АВТОРИЗОВАН + пользователь не имеет доступа
	}
}

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	id := &models.BannerID{}
	var err error
	vars := mux.Vars(r)
	idStr := vars["id"]
	if id.BannerId, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		log.Printf("DeleteBanner: couldn't delete the banner: %v", err)
		JSONError(w, models.ErrorMessage{Error: "No id passed"}, http.StatusBadRequest)
		return
	} else {
		if err := h.service.DeleteBanner(id); err != nil {
			log.Printf("DeleteBanner: couldn;t delete the banner: %v", err)
			JSONError(w, models.ErrorMessage{Error: "Internal Server Error"}, http.StatusInternalServerError) //должно быть две ошибки
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func (h *Handler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	input := &models.BannerUpdateRequest{}
	var err error
	err = json.NewDecoder(r.Body).Decode(input)
	vars := mux.Vars(r)
	idStr := vars["id"]
	input.BannerId = new(uint64)
	fmt.Println(input.TagIDs)
	if *input.BannerId, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		log.Printf("DeleteBanner: couldn't find id: %v", err)
		JSONError(w, models.ErrorMessage{Error: "No id passed"}, http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("UpdateBanner: couldn't decode the json: %v", err)
		JSONError(w, models.ErrorMessage{Error: "couldn't decode the json"}, http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateBanner(input); err != nil {
		log.Printf("UpdateBanner: couldn't decode the json: %v", err)
		JSONError(w, models.ErrorMessage{Error: "No results returned"}, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func JSONError(w http.ResponseWriter, err interface{}, code int) { // куда тебя деть
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
