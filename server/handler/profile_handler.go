package handler

import (
	"api-egressos/service"
	"encoding/json"
	"net/http"
)

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	service := service.NewProfileService()
	profiles, err := service.GetProfiles()

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, profiles)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
