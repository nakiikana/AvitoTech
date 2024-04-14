package handler

import (
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
		JSONError(w, models.ErrorMessage{Error: ErrorBannerTagIdNotFound.Error()}, http.StatusBadRequest)
		return
	}
	input.FeatureID, err = strconv.ParseUint(r.URL.Query().Get("feature_id"), 10, 64)
	if err != nil {
		log.Printf("Banner: can't get feature_id: %v", err)
		JSONError(w, models.ErrorMessage{Error: ErrorBannerFeatureIdNotFound.Error()}, http.StatusBadRequest)
		return
	}
	input.UseLastRevision, err = strconv.ParseBool(r.URL.Query().Get("use_last_revision"))
	fmt.Println(input.UseLastRevision)
	if err != nil {
		input.UseLastRevision = false
	}

	if content, err := h.service.FindBanner(input); err != nil && err.Error() == "no rows found" {
		log.Printf("Couldn't find the required banner: %v", err)
		JSONError(w, models.ErrorMessage{Error: ErrorNoRows.Error()}, http.StatusNotFound)
		return
	} else if err != nil && err.Error() == "no access" {
		log.Printf("FindBanner: no access granted")
		JSONError(w, models.ErrorMessage{Error: ErrorAccess.Error()}, http.StatusForbidden)
		return
	} else if err != nil {
		log.Printf("FindBanner: unexpected error: %v", err)
		JSONError(w, models.ErrorMessage{Error: err.Error()}, http.StatusForbidden)
	} else {
		w.Write(content.Content)
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
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		log.Printf("UpdateBanner: couldn't decode the json: %v", err)
		JSONError(w, models.ErrorMessage{Error: "couldn't decode the json"}, http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]
	input.BannerId = new(uint64)

	if *input.BannerId, err = strconv.ParseUint(idStr, 10, 64); err != nil {
		log.Printf("DeleteBanner: couldn't find id: %v", err)
		JSONError(w, models.ErrorMessage{Error: "No id passed"}, http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateBanner(input); err != nil {
		log.Printf("UpdateBanner: couldn't update the banner: %v", err)
		JSONError(w, models.ErrorMessage{Error: "No results returned"}, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) GetFilteredBanner(w http.ResponseWriter, r *http.Request) {
	input := models.BannerGetAdminRequest{}

	if val, _ := strconv.ParseUint(r.URL.Query().Get("tag_id"), 10, 64); val != 0 {
		input.TagID = new(uint64)
		*input.TagID = val
	}
	if val, _ := strconv.ParseUint(r.URL.Query().Get("feature_id"), 10, 64); val != 0 {
		input.FeatureID = new(uint64)
		*input.FeatureID = val
	}
	if val, _ := strconv.Atoi(r.URL.Query().Get("limit")); val != 0 {
		input.Limit = new(int)
		*input.Limit = val
	}
	if val, _ := strconv.Atoi(r.URL.Query().Get("offset")); val != 0 {
		input.Offset = new(int)
		*input.Offset = val
	}

	bnr, err := h.service.GetFilteredBanner(&input)

	if err != nil {
		log.Printf("GetFilteredBanner: couldn't implement the query: %v", err)
		JSONError(w, models.ErrorMessage{Error: "Couldn't implement the query"}, http.StatusInternalServerError)
		return
	}
	ans, err := json.Marshal(bnr)
	if err != nil {
		log.Printf("GetFilteredBanner: couldn't unmarshall json")
		JSONError(w, models.ErrorMessage{Error: "Couldn't unmarshall json"}, http.StatusInternalServerError)
	}
	w.Write(ans)
}
