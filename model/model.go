package model

import "time"

type Hotel struct {
	RoomId       int           `json:"room_id"`
	HotelName    string        `json:"hotel_name"`
	RatePerNight float64       `json:"rate_per_night"`
	MaxNumGuests int           `json:"max_guests"`
	IsAvailable  bool          `json:"is_available"`
	AvailableDates []AvailableDate `json:"available_dates"` 
}

type AvailableDate struct {
	HotelId       int       `json:"-"` 
	AvailableDate time.Time `json:"available_date"`
}
  