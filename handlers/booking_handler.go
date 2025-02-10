package handlers

import (
	"encoding/json"
	"glofox/models"
	"glofox/storage"
	"net/http"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bookingRequest models.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if bookingRequest.Name == "" || bookingRequest.Date.IsZero() {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	bookingResponse, err := storage.BookClass(bookingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookingResponse)
}
