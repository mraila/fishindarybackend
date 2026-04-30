package handler

import (
	"encoding/json"
	"net/http"

	"fishindary/model"
	"fishindary/service"
)

func CreateCatch(w http.ResponseWriter, r *http.Request) {
	var c model.Catch

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result := service.CreateCatch(c)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetCatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetCatches())
}
