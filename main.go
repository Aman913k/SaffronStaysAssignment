package main

import (
	"github.com/Aman913k/SaffronStaysAssignment/controller"
	"github.com/Aman913k/SaffronStaysAssignment/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB("user=postgres password=securePass123 host=localhost port=5432 dbname=hotels sslmode=disable")

	r := gin.Default()

	r.POST("/hotel/create", controller.CreateHotel)
	r.GET("/hotel/:id", controller.GetHotelDetailsById)

	r.Run(":8000")
}
