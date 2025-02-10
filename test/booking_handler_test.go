package test

import (
	"bytes"
	"encoding/json"
	"glofox/handlers"
	"glofox/models"
	"glofox/storage"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	// Add a class for testing
	startDate, _ := time.Parse(time.RFC3339, "2025-02-11T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2025-02-21T00:00:00Z")
	storage.CreateClass(models.Class{Name: "Pilates", StartDate: startDate, EndDate: endDate, Capacity: 10})

	payload := []byte(`{"name": "Prakash Anand", "date": "2025-02-16T00:00:00Z"}`)
	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateBooking)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response models.BookingResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Member != "Prakash Anand" {
		t.Errorf("Expected member name 'Prakash Anand', got '%v'", response.Member)
	}
}
