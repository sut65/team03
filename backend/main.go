package main

import (
	booking "github.com/sut65/team03/controller/Booking"
	controller "github.com/sut65/team03/controller/Manage_Employee"

	"github.com/sut65/team03/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Officer Routes
	r.GET("/Officers", controller.ListOfficers)
	r.GET("/Officer/:id", controller.GetOfficer)
	r.POST("/Officers", controller.CreateOfficer)
	r.PATCH("/Officers", controller.UpdateOfficer)
	r.DELETE("/Officers/:id", controller.DeleteOfficer)

	r.GET("/Departments", controller.ListDepartments)
	r.GET("/Department/:id", controller.GetDepartment)
	r.POST("/Departments", controller.CreateDepartment)
	r.PATCH("/Departments", controller.UpdateDepartment)
	r.DELETE("/Departments/:id", controller.DeleteDepartment)

	r.GET("/Positions", controller.ListPositions)
	r.GET("/Position/:id", controller.GetPosition)
	r.POST("/Positions", controller.CreatePosition)
	r.PATCH("/Positions", controller.UpdatePosition)
	r.DELETE("/Positions/:id", controller.DeletePosition)

	r.GET("/Locations", controller.ListLocations)
	r.GET("/Location/:id", controller.GetLocation)
	r.POST("/Locations", controller.CreateLocation)
	r.PATCH("/Locations", controller.UpdateLocation)
	r.DELETE("/Locations/:id", controller.DeleteLocation)

	r.GET("/Employees", controller.ListEmployees)
	r.GET("/Employee/:id", controller.GetEmployee)
	r.POST("/Employees", controller.CreateEmployee)
	r.PATCH("/Employees", controller.UpdateEmployee)
	r.DELETE("/Employees/:id", controller.DeleteEmployee)

	//=================================================== Booking Routes
	r.GET("/bookings", booking.ListBookings)
	r.GET("/booking/:id", booking.GetBooking)
	r.GET("/bookings/user/:id", booking.ListBookingsByUID)
	r.POST("/bookings", booking.CreateBooking)
	r.PATCH("/bookings", booking.UpdateBooking)
	r.DELETE("/bookings/:id", booking.DeleteBooking)
	// ---Branch---
	r.GET("/branchs", booking.ListBranchs)
	r.GET("/branch/:id", booking.GetBranch)
	r.POST("/branchs", booking.CreateBranch)
	r.PATCH("/branchs", booking.UpdateBranch)
	r.DELETE("/branchs/:id", booking.DeleteBranch)
	//=================================================== Booking Routes

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
