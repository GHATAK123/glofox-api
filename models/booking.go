package models

import "time"

type BookingRequest struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type BookingResponse struct {
	ClassName string `json:"class_name"`
	Member    string `json:"member"`
	Date      string `json:"date"`
}
