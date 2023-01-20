package main

import (
	booking "github.com/sut65/team03/controller/Booking"
	chk_payment "github.com/sut65/team03/controller/Chk_Payment"
	controller "github.com/sut65/team03/controller/Manage_Employee"
	customer "github.com/sut65/team03/controller/Customer"

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
	//=================================================== Check Payment Routes
	r.GET("/chk_payments", chk_payment.ListCHK_Payments)
	r.GET("/chk_payment/:id", chk_payment.GetCHK_Payment)
	r.POST("/chk_payments", chk_payment.CreateCHK_Payment)
	r.PATCH("/chk_payments", chk_payment.UpdateCHK_Payment)
	r.DELETE("/chk_payments/:id", chk_payment.DeleteCHK_Payment)
	// ---Status---
	r.GET("/chk_payment/statuses", chk_payment.ListCHK_Payments)
	r.GET("/chk_payment/status/:id", chk_payment.GetCHK_Payment)
	r.POST("/chk_payment/statuses", chk_payment.CreateCHK_Payment)
	r.PATCH("/chk_payment/statuses", chk_payment.UpdateCHK_Payment)
	r.DELETE("/chk_payment/statuses/:id", chk_payment.DeleteCHK_Payment)
	//=================================================== Check Payment Routes

	//==================================================Customer Routes
	r.GET("/customers", customer.ListCustomers)
	r.GET("/customer/:id", customer.GetCustomerByID)
	r.POST("/customers",  customer.CreateCustomer)
	r.PATCH("/customers", customer.UpdateCustomer)
	r.DELETE("/customers/:id", customer.DeleteCustomer)
	//Gender
	r.GET("/customers/genders", customer.ListGender)
	r.GET("/customer/genders/:id", customer.GetGender)
	r.POST("/customers/genders", customer.CreateGender)
	r.PATCH("/customers/genders", customer.UpdateGender)
	r.DELETE("/customers/genders/:id",customer.DeleteGender)
	//Memberlevel
	r.GET("/memberlevels", customer.ListMemberlevel)
	r.GET("/customer/memberlevels/:id", customer.GetMemberlevel)
	r.POST("/customers/memberlevels", customer.CreateMemberlevel)
	r.PATCH("/customers/memberlevels", customer.UpdateMemberlevel)
	r.DELETE("/customers/memberlevels/:id",customer.DeleteMemberlevel)
	//Province
	r.GET("/provinces", customer.ListProvince)
	r.GET("/customer/provinces/:id", customer.GetProvince)
	r.POST("/customers/provinces", customer.CreateProvince)
	r.PATCH("/customers/provinces", customer.UpdateProvince)
	r.DELETE("/customers/provinces/:id",customer.DeleteProvince)
	//==================================================Customer Routes

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
