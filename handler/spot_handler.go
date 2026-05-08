package handler

import (
	"encoding/json"
	"fishindary/model"
	"fishindary/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type SpotHandler struct {
	SpotService  *service.SpotService
	CatchService *service.CatchService
}

func NewSpotHandler(s *service.SpotService, c *service.CatchService) *SpotHandler {
	return &SpotHandler{
		SpotService:  s,
		CatchService: c,
	}
}

func validateSpot(s model.Spot) error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}
	if s.BodyOfWater == "" {
		return fmt.Errorf("body of water is required")
	}
	if s.Location.Latitude == 0 && s.Location.Longitude == 0 {
		return fmt.Errorf("location is required")
	}
	return nil
}

func (h *SpotHandler) CreateSpot(w http.ResponseWriter, r *http.Request) {
	var s model.Spot

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := validateSpot(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.SpotService.CreateSpot(s)
	json.NewEncoder(w).Encode(result)
}

func (h *SpotHandler) GetSpots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.SpotService.GetSpots())
}

func (h *SpotHandler) GetCatchesBySpotID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/spots/")
	idStr = strings.TrimSuffix(idStr, "/catches")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid spot ID", http.StatusBadRequest)
		return
	}

	catches := h.CatchService.GetCatchesBySpotID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catches)
}
