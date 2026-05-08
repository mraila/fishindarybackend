package main

import (
	"encoding/json"
	"fishindary/handler"
	"fishindary/model"
	"fishindary/service"
	"fishindary/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCatch(t *testing.T) {

	// Create a new HTTP request to the /catches endpoint
	reqBody := `{
		"fish_type": "pike",
		"weight": 1500,
		"length": 58,
		"lure": "aglia #3",
		"location": {
			"latitude": 60.1695,
			"longitude": 24.9354
		}
	}`
	req := httptest.NewRequest(http.MethodPost, "/catches", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	h := handler.NewCatchHandler(catchSvc)
	h.CreateCatch(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200, got %d", res.StatusCode)
	}

	var response model.Catch
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("Expected ID 1, got %d", response.ID)
	}

	if response.FishType != "pike" {
		t.Errorf("Expected fish type pike, got %s", response.FishType)
	}

	if response.Location.Latitude != 60.1695 {
		t.Errorf("expected lat 60.1695, got %f", response.Location.Latitude)
	}

	if response.Location.Longitude != 24.9354 {
		t.Errorf("expected lon 24.9354, got %f", response.Location.Longitude)
	}
}

func TestCreateCatch_InvalidInput(t *testing.T) {
	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	h := handler.NewCatchHandler(catchSvc)

	body := `{
		"fish_type": "",
		"weight": 0,
		"length": 0,
		"location": { "lat": 0, "lng": 0 }
	}`

	req := httptest.NewRequest(http.MethodPost, "/catches", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.CreateCatch(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}

func TestGetCatches(t *testing.T) {
	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	h := handler.NewCatchHandler(catchSvc)

	// seed data first
	_ = catchSvc.CreateCatch(model.Catch{
		FishType: "pike",
		Weight:   1500,
		Length:   58,
		Lure:     "aglia #3",
		Location: model.Location{Latitude: 60.1695, Longitude: 24.9354},
	})

	req := httptest.NewRequest(http.MethodGet, "/catches", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(h.GetCatches).ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200")
	}

	var result []model.Catch
	json.NewDecoder(res.Body).Decode(&result)

	if len(result) != 1 {
		t.Fatalf("expected 1 catch, got %d", len(result))
	}
}
