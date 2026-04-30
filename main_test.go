package main

import (
	"encoding/json"
	"fishindary/handler"
	"fishindary/model"
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

	httpHandler := http.HandlerFunc(handler.CreateCatch)
	httpHandler.ServeHTTP(w, req)

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

func TestGetCatches(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/catches", nil)

	w := httptest.NewRecorder()

	httpHandler := http.HandlerFunc(handler.GetCatches)
	httpHandler.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200")
	}

	var result []model.Catch
	json.NewDecoder(res.Body).Decode(&result)

	if len(result) == 0 {
		t.Fatalf("expected at least 1 catch")
	}
}
