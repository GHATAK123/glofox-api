package storage

import (
	"errors"
	"glofox/models"
)

var (
	classes  []models.Class
	bookings []models.BookingResponse
)

func CreateClass(newClass models.Class) {
	classes = append(classes, newClass)
}

func BookClass(bookingRequest models.BookingRequest) (models.BookingResponse, error) {
	for _, class := range classes {
		if !bookingRequest.Date.Before(class.StartDate) && !bookingRequest.Date.After(class.EndDate) {
			bookingResponse := models.BookingResponse{
				ClassName: class.Name,
				Member:    bookingRequest.Name,
				Date:      bookingRequest.Date.Format("2006-01-02"),
			}
			bookings = append(bookings, bookingResponse)
			return bookingResponse, nil
		}
	}
	return models.BookingResponse{}, errors.New("class not available on the requested date")
}

func GetClasses() []models.Class {
	return classes
}

func GetBookings() []models.BookingResponse {
	return bookings
}
