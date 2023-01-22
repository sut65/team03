package main

import (
	booking "github.com/sut65/team03/controller/Booking"
	check "github.com/sut65/team03/controller/CheckInOut"
	chk_payment "github.com/sut65/team03/controller/Chk_Payment"
	customer "github.com/sut65/team03/controller/Customer"
	controller "github.com/sut65/team03/controller/Manage_Employee"
	payment "github.com/sut65/team03/controller/Payment"
	repreq "github.com/sut65/team03/controller/RepReq"
	reviewht "github.com/sut65/team03/controller/Review"
	room "github.com/sut65/team03/controller/Room"
	service "github.com/sut65/team03/controller/Service"

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
	r.POST("/customers", customer.CreateCustomer)
	r.PATCH("/customers", customer.UpdateCustomer)
	r.DELETE("/customers/:id", customer.DeleteCustomer)
	//Gender
	r.GET("/customers/genders", customer.ListGender)
	r.GET("/customer/genders/:id", customer.GetGender)
	r.POST("/customers/genders", customer.CreateGender)
	r.PATCH("/customers/genders", customer.UpdateGender)
	r.DELETE("/customers/genders/:id", customer.DeleteGender)
	//Memberlevel
	r.GET("/memberlevels", customer.ListMemberlevel)
	r.GET("/customer/memberlevels/:id", customer.GetMemberlevel)
	r.POST("/customers/memberlevels", customer.CreateMemberlevel)
	r.PATCH("/customers/memberlevels", customer.UpdateMemberlevel)
	r.DELETE("/customers/memberlevels/:id", customer.DeleteMemberlevel)
	//Province
	r.GET("/provinces", customer.ListProvince)
	r.GET("/customer/provinces/:id", customer.GetProvince)
	r.POST("/customers/provinces", customer.CreateProvince)
	r.PATCH("/customers/provinces", customer.UpdateProvince)
	r.DELETE("/customers/provinces/:id", customer.DeleteProvince)
	//==================================================Customer Routes

	//========================= checkInOut routes
	//status
	r.GET("/checkinoutstatus/:id", check.GetCheckInOutStatus)
	r.GET("/checkinoutstatuses", check.ListCheckInOutStatuses)
	r.POST("/checkinoutstatus", check.CreateCheckInOutStatus)
	r.PATCH("/checkinoutstatus", check.UpdateCheckInOutStatus)
	r.DELETE("/checkinoutstatus/:id", check.DeleteCheckInOutStatus)
	//main
	r.GET("/checkinout/:id", check.GetCheckInOut)
	r.GET("/checkinouts", check.ListCheckInOuts)
	r.POST("/checkinout", check.CreateCheckInOut)
	r.PATCH("/checkinout", check.UpdateCheckInOut)
	r.DELETE("/checkinout/:id", check.DeleteCheckInOut)

	//========================= repreq routes
	//type
	r.GET("/repairtype/:id", repreq.GetRepairType)
	r.GET("/repairtypes", repreq.ListRepairTypes)
	r.POST("/repairtype", repreq.CreateRepairType)
	r.PATCH("/repairtype", repreq.UpdateRepairType)
	r.DELETE("/repairtype/:id", repreq.DeleteRepairType)
	//status
	r.GET("/repairstatus/:id", repreq.GetRepairStatus)
	r.GET("/repairstatuses", repreq.ListRepairStatuses)
	r.POST("/repairstatus", repreq.CreateRepairStatus)
	r.PATCH("/repairstatus", repreq.UpdateRepairStatus)
	r.DELETE("/repairstatus/:id", repreq.DeleteRepairStatus)
	//main
	r.GET("/repairreq/:id", repreq.GetRepairReq)
	r.GET("/repairreqs", repreq.ListRepairReqs)
	r.POST("/repairreq", repreq.CreateRepairReq)
	r.PATCH("/repairreq", repreq.UpdateRepairReq)
	r.DELETE("/repairreq/:id", repreq.DeleteRepairReq)

	//==================================================Room Routes
	r.GET("/Rooms", room.ListRooms)
	r.GET("/Room/:id", room.GetRoom)
	r.POST("/Rooms", room.CreateRoom)
	r.PATCH("/Rooms", room.UpdateRoom)
	r.DELETE("/Rooms/:id", room.DeleteRoom)

	r.GET("/RoomTypes", room.ListRoomTypes)
	r.GET("/RoomType/:id", room.GetRoomType)
	r.POST("/RoomTypes", room.CreateRoomType)
	r.PATCH("/RoomTypes", room.UpdateRoomType)
	r.DELETE("/RoomTypes/:id", room.DeleteRoomType)

	r.GET("/RoomZones", room.ListRoomZones)
	r.GET("/RoomZone/:id", room.GetRoomZone)
	r.POST("/RoomZones", room.CreateRoomZone)
	r.PATCH("/RoomZones", room.UpdateRoomZone)
	r.DELETE("/RoomZones/:id", room.DeleteRoomZone)

	r.GET("/States", room.ListStates)
	r.GET("/State/:id", room.GetState)
	r.POST("/States", room.CreateState)
	r.PATCH("/States", room.UpdateState)
	r.DELETE("/States/:id", room.DeleteState)
	//===================================================Room

	// ======================================= PAYMENT
	r.GET("/payment", payment.ListPayments)
	r.GET("/payment/:id", payment.GetPayment)
	r.GET("/payment/user/:id", payment.ListPaymentByUID)
	r.POST("/payment", payment.CreatePayment)
	// ======================================= PAYMENT

	// ======================================= SERVICE
	r.GET("/services", service.ListServices)
	r.GET("/service/:id", service.GetService)
	r.GET("/services/user/:id", service.ListServicesByUID)
	r.POST("/service", service.CreateService)
	r.PATCH("/services", service.UpdateService)
	r.DELETE("/services/:id", service.DeleteService)
	// ======================================= SERVICE

	// Run the server

	//----------review----------------------
	// Review Routes
	r.GET("/Reviews", reviewht.ListReviews)
	r.GET("/Review/:id", reviewht.GetReview)
	r.POST("/Reviews", reviewht.CreateReview)
	r.PATCH("/Reviews", reviewht.UpdateReview)
	r.DELETE("/Reviews/:id", reviewht.DeleteReview)

	// Systemwork Routes
	r.GET("/Systemworks", reviewht.ListSystemworks)
	r.GET("/Systemwork/:id", reviewht.GetSystemwork)
	r.POST("/Systemworks", reviewht.CreateSystemwork)
	r.PATCH("/Systemworks", reviewht.UpdateSystemwork)
	r.DELETE("/Systemworks/:id", reviewht.DeleteSystemwork)

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
