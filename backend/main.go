package main

import (
	"github.com/sut65/team03/controller"
	booking "github.com/sut65/team03/controller/Booking"
	check "github.com/sut65/team03/controller/CheckInOut"
	chk_payment "github.com/sut65/team03/controller/Chk_Payment"
	customer "github.com/sut65/team03/controller/Customer"
	employee "github.com/sut65/team03/controller/Manage_Employee"
	payment "github.com/sut65/team03/controller/Payment"
	repreq "github.com/sut65/team03/controller/RepReq"
	reviewht "github.com/sut65/team03/controller/Review"
	room "github.com/sut65/team03/controller/Room"
	service "github.com/sut65/team03/controller/Service"
	"github.com/sut65/team03/middlewares"

	"github.com/sut65/team03/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// Officer Routes
			protected.GET("/Officers", employee.ListOfficers)
			protected.GET("/Officer/:id", employee.GetOfficer)
			protected.POST("/Officers", employee.CreateOfficer)

			protected.GET("/Departments", employee.ListDepartments)
			protected.GET("/Department/:id", employee.GetDepartment)
			protected.POST("/Departments", employee.CreateDepartment)

			protected.GET("/Positions", employee.ListPositions)
			protected.GET("/Position/:id", employee.GetPosition)
			protected.POST("/Positions", employee.CreatePosition)

			protected.GET("/Locations", employee.ListLocations)
			protected.GET("/Location/:id", employee.GetLocation)
			protected.POST("/Locations", employee.CreateLocation)
			protected.PATCH("/Locations", employee.UpdateLocation)
			protected.DELETE("/Locations/:id", employee.DeleteLocation)

			protected.GET("/Employees", employee.ListEmployees)
			protected.GET("/Employee/:id", employee.GetEmployee)
			protected.POST("/Employees", employee.CreateEmployee)
			protected.PATCH("/Employees", employee.UpdateEmployee)
			protected.DELETE("/Employees/:id", employee.DeleteEmployee)

			//=================================================== Booking Routes
			protected.GET("/bookings", booking.ListBookings)
			protected.GET("/booking/:id", booking.GetBooking)
			protected.GET("/bookings/user/:id", booking.ListBookingsByUID)
			protected.POST("/bookings", booking.CreateBooking)
			protected.PATCH("/bookings", booking.UpdateBooking)
			protected.DELETE("/bookings/:id", booking.DeleteBooking)
			// ---Branch---
			protected.GET("/branchs", booking.ListBranchs)
			protected.GET("/branch/:id", booking.GetBranch)
			protected.POST("/branchs", booking.CreateBranch)
			protected.PATCH("/branchs", booking.UpdateBranch)
			protected.DELETE("/branchs/:id", booking.DeleteBranch)
			//=================================================== Booking Routes
			//=================================================== Check Payment Routes
			protected.GET("/chk_payments", chk_payment.ListCHK_Payments)
			protected.GET("/chk_payment/:id", chk_payment.GetCHK_Payment)
			protected.POST("/chk_payments", chk_payment.CreateCHK_Payment)
			protected.PATCH("/chk_payments", chk_payment.UpdateCHK_Payment)
			protected.DELETE("/chk_payments/:id", chk_payment.DeleteCHK_Payment)
			// ---Status---
			protected.GET("/chk_payment/statuses", chk_payment.ListStatuses)
			protected.GET("/chk_payment/status/:id", chk_payment.GetStatus)
			protected.POST("/chk_payment/statuses", chk_payment.CreateStatus)
			protected.PATCH("/chk_payment/statuses", chk_payment.UpdateStatus)
			protected.DELETE("/chk_payment/statuses/:id", chk_payment.DeleteStatus)
			//=================================================== Check Payment Routes

			//==================================================Customer Routes
			protected.GET("/customers", customer.ListCustomers)
			protected.GET("/customer/:id", customer.GetCustomerByID)
			protected.POST("/customers", customer.CreateCustomer)
			protected.PATCH("/customers", customer.UpdateCustomer)
			protected.DELETE("/customers/:id", customer.DeleteCustomer)
			//Gender
			protected.GET("/customers/genders", customer.ListGender)
			protected.GET("/customer/genders/:id", customer.GetGender)
			protected.POST("/customers/genders", customer.CreateGender)
			protected.PATCH("/customers/genders", customer.UpdateGender)
			protected.DELETE("/customers/genders/:id", customer.DeleteGender)
			//Memberlevel
			protected.GET("/memberlevels", customer.ListMemberlevel)
			protected.GET("/customer/memberlevels/:id", customer.GetMemberlevel)
			protected.POST("/customers/memberlevels", customer.CreateMemberlevel)
			protected.PATCH("/customers/memberlevels", customer.UpdateMemberlevel)
			protected.DELETE("/customers/memberlevels/:id", customer.DeleteMemberlevel)
			//Province
			protected.GET("/provinces", customer.ListProvince)
			protected.GET("/customer/provinces/:id", customer.GetProvince)
			protected.POST("/customers/provinces", customer.CreateProvince)
			protected.PATCH("/customers/provinces", customer.UpdateProvince)
			protected.DELETE("/customers/provinces/:id", customer.DeleteProvince)
			//==================================================Customer Routes

			//========================= checkInOut routes
			//status
			protected.GET("/checkinoutstatus/:id", check.GetCheckInOutStatus)
			protected.GET("/checkinoutstatuses", check.ListCheckInOutStatuses)
			protected.POST("/checkinoutstatus", check.CreateCheckInOutStatus)
			protected.PATCH("/checkinoutstatus", check.UpdateCheckInOutStatus)
			protected.DELETE("/checkinoutstatus/:id", check.DeleteCheckInOutStatus)
			//main
			protected.GET("/checkinout/:id", check.GetCheckInOut)
			protected.GET("/checkinouts", check.ListCheckInOuts)
			protected.POST("/checkinout", check.CreateCheckInOut)
			protected.PATCH("/checkinout", check.UpdateCheckInOut)
			protected.DELETE("/checkinout/:id", check.DeleteCheckInOut)
			protected.PATCH("/checkinout/:id", check.CheckOut)

			//========================= repreq routes
			//type
			protected.GET("/repairtype/:id", repreq.GetRepairType)
			protected.GET("/repairtypes", repreq.ListRepairTypes)
			protected.POST("/repairtype", repreq.CreateRepairType)
			protected.PATCH("/repairtype", repreq.UpdateRepairType)
			protected.DELETE("/repairtype/:id", repreq.DeleteRepairType)
			//status
			protected.GET("/repairstatus/:id", repreq.GetRepairStatus)
			protected.GET("/repairstatuses", repreq.ListRepairStatuses)
			protected.POST("/repairstatus", repreq.CreateRepairStatus)
			protected.PATCH("/repairstatus", repreq.UpdateRepairStatus)
			protected.DELETE("/repairstatus/:id", repreq.DeleteRepairStatus)
			//main
			protected.GET("/repairreq/:id", repreq.GetRepairReq)
			protected.GET("/repairreqs", repreq.ListRepairReqs)
			protected.POST("/repairreq", repreq.CreateRepairReq)
			protected.PATCH("/repairreq", repreq.UpdateRepairReq)
			protected.DELETE("/repairreq/:id", repreq.DeleteRepairReq)

			//==================================================Room Routes
			protected.GET("/Rooms", room.ListRooms)
			protected.GET("/Room/:id", room.GetRoom)
			protected.POST("/Rooms", room.CreateRoom)
			protected.PATCH("/Rooms", room.UpdateRoom)
			protected.DELETE("/Rooms/:id", room.DeleteRoom)

			protected.GET("/RoomTypes", room.ListRoomTypes)
			protected.GET("/RoomType/:id", room.GetRoomType)
			protected.POST("/RoomTypes", room.CreateRoomType)
			protected.PATCH("/RoomTypes", room.UpdateRoomType)
			protected.DELETE("/RoomTypes/:id", room.DeleteRoomType)

			protected.GET("/RoomZones", room.ListRoomZones)
			protected.GET("/RoomZone/:id", room.GetRoomZone)
			protected.POST("/RoomZones", room.CreateRoomZone)
			protected.PATCH("/RoomZones", room.UpdateRoomZone)
			protected.DELETE("/RoomZones/:id", room.DeleteRoomZone)

			protected.GET("/States", room.ListStates)
			protected.GET("/State/:id", room.GetState)
			protected.POST("/States", room.CreateState)
			protected.PATCH("/States", room.UpdateState)
			protected.DELETE("/States/:id", room.DeleteState)
			//===================================================Room

			// ======================================= PAYMENT
			protected.GET("/payments", payment.ListPayments)
			protected.GET("/payment/:id", payment.GetPayment)
			protected.GET("/payment/customer/:id", payment.ListPaymentByUID)
			protected.POST("/payment", payment.CreatePayment)
			// ======================================= PAYMENT

			// ======================================= SERVICE
			protected.GET("/services", service.ListServices)
			protected.GET("/service/:id", service.GetService)
			protected.GET("/services/customer/:id", service.ListServicesByUID)
			protected.POST("/service", service.CreateService)
			protected.PATCH("/services", service.UpdateService)
			protected.DELETE("/services/:id", service.DeleteService)
			// ======================================= SERVICE

			// Run the server

			//----------review----------------------
			// Review Routes
			protected.GET("/Reviews", reviewht.ListReviews)
			protected.GET("/Review/:id", reviewht.GetReview)
			protected.POST("/Reviews", reviewht.CreateReview)
			protected.PATCH("/Reviews", reviewht.UpdateReview)
			protected.DELETE("/Reviews/:id", reviewht.DeleteReview)

			// Systemwork Routes
			protected.GET("/Systemworks", reviewht.ListSystemworks)
			protected.GET("/Systemwork/:id", reviewht.GetSystemwork)
			protected.POST("/Systemworks", reviewht.CreateSystemwork)
			protected.PATCH("/Systemworks", reviewht.UpdateSystemwork)
			protected.DELETE("/Systemworks/:id", reviewht.DeleteSystemwork)
		}
	}
	r.POST("/login", controller.Login)

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
