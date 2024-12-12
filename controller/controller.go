package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Aman913k/SaffronStaysAssignment/database"
	"github.com/Aman913k/SaffronStaysAssignment/model"
	"github.com/gin-gonic/gin"
)

func CreateHotel(c *gin.Context) {
	var hotel model.Hotel

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	query := `INSERT INTO hotels (hotel_name, rate_per_night, max_guests, is_available) 
	          VALUES ($1, $2, $3, $4) RETURNING room_id`
	err := database.DB.QueryRow(query, hotel.HotelName, hotel.RatePerNight, hotel.MaxNumGuests, hotel.IsAvailable).Scan(&hotel.RoomId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		return
	}

	if hotel.IsAvailable && len(hotel.AvailableDates) > 0 {
		for _, availableDate := range hotel.AvailableDates {
			query := `INSERT INTO available_dates (hotel_id, available_date) VALUES ($1, $2)`
			_, err := database.DB.Exec(query, hotel.RoomId, availableDate.AvailableDate)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert available dates"})
				return
			}
		}
	} else if hotel.IsAvailable {
		populateAvailableDates(hotel.RoomId)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hotel created successfully", "hotel": hotel})
}

func populateAvailableDates(hotelID int) {
	now := time.Now()
	for i := 0; i < 150; i++ {
		date := now.AddDate(0, 0, i)
		query := `INSERT INTO available_dates (hotel_id, available_date) VALUES ($1, $2)`
		_, err := database.DB.Exec(query, hotelID, date)
		if err != nil {
			panic(err)
		}
	}
}

func GetHotelDetailsById(c *gin.Context) {
	idParam := c.Param("id")
	hotelID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	var hotel model.Hotel
	query := `SELECT room_id, hotel_name, rate_per_night, max_guests, is_available 
			  FROM hotels WHERE room_id = $1`
	err = database.DB.QueryRow(query, hotelID).Scan(&hotel.RoomId, &hotel.HotelName, &hotel.RatePerNight, &hotel.MaxNumGuests, &hotel.IsAvailable)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	availableDatesQuery := `SELECT available_date FROM available_dates WHERE hotel_id = $1`
	rows, err := database.DB.Query(availableDatesQuery, hotelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available dates"})
		return
	}
	defer rows.Close()

	var availableDates []model.AvailableDate
	for rows.Next() {
		var availableDate model.AvailableDate
		err := rows.Scan(&availableDate.AvailableDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse available date"})
			return
		}
		availableDates = append(availableDates, availableDate)
	}

	hotel.AvailableDates = availableDates

	occupancy, rates := calculateStats(hotelID)

	c.JSON(http.StatusOK, gin.H{
		"hotel":     hotel,
		"occupancy": occupancy,
		"rates":     rates,
	})
}

func calculateStats(hotelID int) (map[string]float64, map[string]float64) {

	occupancyQuery := `SELECT COUNT(*) FROM available_dates WHERE hotel_id = $1 AND available_date > CURRENT_DATE`
	var availableDays int
	database.DB.QueryRow(occupancyQuery, hotelID).Scan(&availableDays)

	totalDays := 150
	occupancy := float64(availableDays) / float64(totalDays) * 100

	rateQuery := `SELECT MAX(rate_per_night), MIN(rate_per_night), AVG(rate_per_night) FROM hotels WHERE room_id = $1`
	var maxRate, minRate, avgRate float64
	database.DB.QueryRow(rateQuery, hotelID).Scan(&maxRate, &minRate, &avgRate)

	return map[string]float64{"percentage": occupancy}, map[string]float64{
		"average": avgRate,
		"highest": maxRate,
		"lowest":  minRate,
	}
}
