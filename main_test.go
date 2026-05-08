package main

import (
	"encoding/json"
	"fishindary/handler"
	"fishindary/model"
	"fishindary/service"
	"fishindary/store"
	"fmt"
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

func TestCreateCatch_WithValidSpot(t *testing.T) {
	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	catchHandler := handler.NewCatchHandler(catchSvc)

	// create spot first
	spot := spotSvc.CreateSpot(model.Spot{
		Name: "Lake edge",
		Location: model.Location{
			Latitude:  54.7,
			Longitude: 25.3,
		},
	})

	body := fmt.Sprintf(`{
		"fish_type": "pike",
		"weight": 2000,
		"length": 70,
		"lure": "spinner",
		"spot_id": %d,
		"location": {
			"latitude": 54.7,
			"longitude": 25.3
		}
	}`, spot.ID)

	req := httptest.NewRequest(http.MethodPost, "/catches", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	catchHandler.CreateCatch(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}

	var response model.Catch
	json.NewDecoder(res.Body).Decode(&response)

	if response.SpotID == nil {
		t.Fatalf("expected spot id to be set")
	}

	if *response.SpotID != spot.ID {
		t.Fatalf("expected spot id %d, got %d", spot.ID, *response.SpotID)
	}
}

func TestCreateCatch_InvalidSpotID(t *testing.T) {
	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	catchHandler := handler.NewCatchHandler(catchSvc)

	body := `{
		"fish_type": "pike",
		"weight": 1000,
		"length": 50,
		"lure": "spinner",
		"spot_id": 999,
		"location": {
			"latitude": 54.7,
			"longitude": 25.3
		}
	}`

	req := httptest.NewRequest(http.MethodPost, "/catches", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	catchHandler.CreateCatch(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", res.StatusCode)
	}
}

func TestGetCatchesBySpotID(t *testing.T) {
	st := store.NewStore()

	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	spotHandler := handler.NewSpotHandler(spotSvc, catchSvc)

	// create spot
	spot := spotSvc.CreateSpot(model.Spot{
		Name: "Perch bay",
		Location: model.Location{
			Latitude:  55,
			Longitude: 25,
		},
	})

	// create catch linked to spot
	catchSvc.CreateCatch(model.Catch{
		FishType: "perch",
		Weight:   400,
		Length:   30,
		Lure:     "jig",
		SpotID:   &spot.ID,
		Location: model.Location{
			Latitude:  55,
			Longitude: 25,
		},
	})

	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/spots/%d/catches", spot.ID),
		nil,
	)

	w := httptest.NewRecorder()

	spotHandler.GetCatchesBySpotID(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}

	var result []model.Catch

	err := json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 catch, got %d", len(result))
	}

	if result[0].FishType != "perch" {
		t.Fatalf("expected perch, got %s", result[0].FishType)
	}
}
