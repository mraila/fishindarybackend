package handler

import (
	"encoding/json"
	"fishindary/model"
	"fishindary/service"
	"fmt"
	"net/http"
)

type CatchHandler struct {
	Service *service.CatchService
}

func NewCatchHandler(s *service.CatchService) *CatchHandler {
	return &CatchHandler{Service: s}
}

func (h *CatchHandler) CreateCatch(w http.ResponseWriter, r *http.Request) {
	var c model.Catch

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.validateCatch(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received catch: %+v\n", c)

	result := h.Service.CreateCatch(c)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *CatchHandler) validateCatch(c model.Catch) error {
	if c.FishType == "" {
		return fmt.Errorf("fish type is required")
	}
	if c.Weight <= 0 {
		return fmt.Errorf("weight must be a positive number")
	}
	if c.Length <= 0 {
		return fmt.Errorf("length must be a positive number")
	}
	if c.Lure == "" {
		return fmt.Errorf("lure is required")
	}
	if c.Location.Latitude == 0 && c.Location.Longitude == 0 {
		return fmt.Errorf("location is required")
	}
	if c.SpotID != nil {
		if !h.Service.SpotService.SpotExists(*c.SpotID) {
			return fmt.Errorf("spot does not exist")
		}
	}
	return nil
}

func (h *CatchHandler) GetCatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Service.GetCatches())
}
