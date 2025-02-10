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

func TestCreateClass(t *testing.T) {
	payload := []byte(`{"name": "Zumba", "start_date": "2025-02-11T00:00:00Z", "end_date": "2025-02-21T00:00:00Z", "capacity": 15}`)
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateClass)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response models.Class
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name != "Zumba" {
		t.Errorf("Expected class name 'Zumba', got '%v'", response.Name)
	}
}
