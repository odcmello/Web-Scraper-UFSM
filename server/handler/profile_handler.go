package handler

import (
	"api-egressos/model"
	"api-egressos/service"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"log"
)

// type ProfilePageData struct {
// 	Profiles []model.Profile
// }

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	service := service.NewProfileService()
	filters, err := applyFilters(r)

	if err != nil {
		http.Error(w, "Invalid url parameter", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "filters", filters)
	profiles, err := service.GetProfiles(ctx)

	if err != nil {
		log.Println(err)
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

func applyFilters(r *http.Request) (map[string]interface{}, error) {
	filters := make(map[string]interface{})
	queries := r.URL.Query()
	for key, value := range queries {
		if _, ok := model.FiltersMapper[key]; ok {
			model.FiltersMapper[key](filters, value[0])
		} else {
			return nil, errors.New("invalid url parameter")
		}
	}
	return filters, nil
}

// func GetPage(w http.ResponseWriter, r *http.Request) {
// 	service := service.NewProfileService()
// 	profiles, err := service.GetProfiles()

// 	if err != nil {
// 		http.Error(w, "Internal Error", http.StatusInternalServerError)
// 		return
// 	}

// 	var profileData []model.Profile

// 	for _, value := range profiles {
// 		profile := model.Profile{
// 			ID:        value.ID,
// 			Name:      value.Name,
// 			JobTitle:  value.JobTitle,
// 			AnoEvasao: value.AnoEvasao,
// 			Curso:     value.Curso,
// 			Company:   value.Company,
// 			Location:  value.Location,
// 			URL:       value.Location,
// 		}

// 		profileData = append(profileData, profile)
// 	}

// 	absPath, _ := filepath.Abs("./pages/index.html")

// 	log.Println(absPath)

// 	templ := template.Must(template.ParseFiles(absPath))
// 	data := ProfilePageData{
// 		Profiles: profileData,
// 	}
// 	templ.Execute(w, data)
// }
