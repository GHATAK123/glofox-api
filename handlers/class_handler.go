package handlers

import (
	"encoding/json"
	"glofox/models"
	"glofox/storage"
	"net/http"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {
	var newClass models.Class
	if err := json.NewDecoder(r.Body).Decode(&newClass); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if newClass.Name == "" || newClass.Capacity <= 0 || newClass.StartDate.IsZero() || newClass.EndDate.IsZero() {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	storage.CreateClass(newClass)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClass)
}
