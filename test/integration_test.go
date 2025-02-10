package test

import (
	"bytes"
	"encoding/json"
	"glofox/handlers"
	"glofox/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegration(t *testing.T) {
	// Create a class
	classPayload := []byte(`{"name": "Yoga", "start_date": "2025-02-11T00:00:00Z", "end_date": "2025-02-21T00:00:00Z", "capacity": 15}`)
	classReq, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(classPayload))
	if err != nil {
		t.Fatal(err)
	}

	classRR := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateClass)
	handler.ServeHTTP(classRR, classReq)

	if classRR.Code != http.StatusCreated {
		t.Errorf("CreateClass returned wrong status code: got %v want %v", classRR.Code, http.StatusCreated)
	}

	// Book a class
	bookingPayload := []byte(`{"name": "Prakash Anand", "date": "2025-02-19T00:00:00Z"}`)
	bookingReq, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(bookingPayload))
	if err != nil {
		t.Fatal(err)
	}

	bookingRR := httptest.NewRecorder()
	handler = http.HandlerFunc(handlers.CreateBooking)
	handler.ServeHTTP(bookingRR, bookingReq)

	if bookingRR.Code != http.StatusCreated {
		t.Errorf("CreateBooking returned wrong status code: got %v want %v", bookingRR.Code, http.StatusCreated)
	}

	var bookingResponse models.BookingResponse
	if err := json.NewDecoder(bookingRR.Body).Decode(&bookingResponse); err != nil {
		t.Fatal(err)
	}

	if bookingResponse.Member != "Prakash Anand" {
		t.Errorf("Expected member name 'Prakash Anand', got '%v'", bookingResponse.Member)
	}
}
