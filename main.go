package main

import (
	"glofox/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/classes", handlers.CreateClass)
	http.HandleFunc("/bookings", handlers.CreateBooking)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
