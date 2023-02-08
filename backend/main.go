package main

import (
	"github.com/sut65/team03/controller"
	booking "github.com/sut65/team03/controller/Booking"
	check "github.com/sut65/team03/controller/CheckInOut"
	checkroom "github.com/sut65/team03/controller/Checkroom"
	chk_payment "github.com/sut65/team03/controller/Chk_Payment"
	customer "github.com/sut65/team03/controller/Customer"
	employee "github.com/sut65/team03/controller/Manage_Employee"
	payment "github.com/sut65/team03/controller/Payment"
	repreq "github.com/sut65/team03/controller/RepReq"
	reviewht "github.com/sut65/team03/controller/Review"
	room "github.com/sut65/team03/controller/Room"
	service "github.com/sut65/team03/controller/Service"
	storage "github.com/sut65/team03/controller/Storage"
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

			protected.GET("/Departments", employee.ListDepartments)
			protected.GET("/Department/:id", employee.GetDepartment)

			protected.GET("/Positions", employee.ListPositions)
			protected.GET("/Position/:id", employee.GetPosition)

			protected.GET("/Locations", employee.ListLocations)
			protected.GET("/Location/:id", employee.GetLocation)

			protected.GET("/Employees", employee.ListEmployees)
			protected.GET("/Employee/:id", employee.GetEmployee)
			protected.GET("Employees/officer/:id", employee.ListEmplooyeeByUID)
			protected.POST("/Employees", employee.CreateEmployee)
			protected.PATCH("/Employees", employee.UpdateEmployee)
			protected.DELETE("/Employees/:id", employee.DeleteEmployee)

			//=================================================== Booking Routes
			protected.GET("/bookings", booking.ListBookings)
			protected.GET("/booking/:id", booking.GetBooking)
			protected.GET("/bookings/user/:id", booking.ListBookingsByUID)
			protected.POST("/bookings", booking.CreateBooking)
			protected.PATCH("/bookings/:id", booking.UpdateBooking)
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
			protected.PATCH("/chk_payments/:id", chk_payment.UpdateCHK_Payment)
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
			//Nametitle
			protected.GET("/nametitles", customer.ListNametitle)
			protected.GET("/customer/nametitles/:id", customer.GetNametitle)
			protected.POST("/customers/nametitles", customer.CreateNametitle)
			protected.PATCH("/customers/nametitles", customer.UpdateNametitle)
			protected.DELETE("/customers/nametitles/:id", customer.DeleteNametitle)
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
			protected.PATCH("/checkin", check.UpdateCheckIn)
			protected.PATCH("/checkout", check.UpdateCheckOut)
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
			//main
			protected.GET("/repairreq/:id", repreq.GetRepairReq)
			protected.GET("/repairreqs", repreq.ListRepairReqs)
			protected.POST("/repairreq", repreq.CreateRepairReq)
			protected.PATCH("/repairreq", repreq.UpdateRepairReq)
			protected.DELETE("/repairreq/:id", repreq.DeleteRepairReq)

			//==================================================Room Routes
			protected.GET("/rooms", room.ListRooms)
			protected.GET("/room/:id", room.GetRoom)
			protected.POST("/rooms", room.CreateRoom)
			protected.PATCH("/rooms", room.UpdateRoom)
			protected.DELETE("/rooms/:id", room.DeleteRoom)

			protected.GET("/room_types", room.ListRoomTypes)
			protected.GET("/room_types/:id", room.GetRoomType)
			protected.POST("/room_types", room.CreateRoomType)
			protected.PATCH("/room_types", room.UpdateRoomType)
			protected.DELETE("/room_typss/:id", room.DeleteRoomType)

			protected.GET("/room_zones", room.ListRoomZones)
			protected.GET("/room_zone/:id", room.GetRoomZone)
			protected.POST("/room_zones", room.CreateRoomZone)
			protected.PATCH("/room_zones", room.UpdateRoomZone)
			protected.DELETE("/room_zones/:id", room.DeleteRoomZone)

			protected.GET("/states", room.ListStates)
			protected.GET("/state/:id", room.GetState)
			protected.POST("/states", room.CreateState)
			protected.PATCH("/states", room.UpdateState)
			protected.DELETE("/states/:id", room.DeleteState)
			//===================================================Room

			// ======================================= PAYMENT
			protected.GET("/payments", payment.ListPayments)
			protected.GET("/payment/:id", payment.GetPayment)
			protected.GET("/payment/customer/:id", payment.ListPaymentByUID)
			protected.POST("/payment", payment.CreatePayment)
			protected.PATCH("/payments", payment.UpdatePayment)

			protected.GET("/paymentmethods", payment.ListPaymentMethods)
			protected.GET("/methods/paymet/:id", payment.ListMethodsByPID)
			protected.GET("/method/:id", payment.GetMethod)
			protected.GET("/places", payment.ListPlaces)
			// ======================================= PAYMENT

			// ======================================= SERVICE
			protected.GET("/services", service.ListServices)
			protected.GET("/service/:id", service.GetService)
			protected.GET("/services/customer/:id", service.ListServicesByUID)
			protected.POST("/service", service.CreateService)
			protected.PATCH("/services", service.UpdateService)
			protected.DELETE("/services/:id", service.DeleteService)

			protected.GET("/room/customer/:id", service.GetRoomByCID)
			protected.GET("/foods", service.ListFoods)
			protected.GET("/drinks", service.ListDrinks)
			protected.GET("/accessories", service.ListAccessories)
			// ======================================= SERVICE

			// Run the server

			//----------review----------------------
			// Review Routes
			r.GET("/Reviews", reviewht.ListReviews)
			protected.GET("/Review/:id", reviewht.GetReview)
			protected.POST("/Reviews", reviewht.CreateReview)
			protected.PATCH("/Reviews", reviewht.UpdateReview)
			protected.DELETE("/Reviews/:id", reviewht.DeleteReview)

			// Systemwork Routes
			protected.GET("/Systemworks", reviewht.ListSystemworks)

			//==================================================Storage Routes
			protected.GET("/storages", storage.ListStorages)
			protected.GET("/storage/:id", storage.GetStorage)
			protected.POST("/storages", storage.CreateStorage)
			protected.PATCH("/storages", storage.UpdateStorage)
			protected.DELETE("/storages/:id", storage.DeleteStorage)

			protected.GET("/products", storage.ListProducts)
			protected.GET("/product/:id", storage.GetProduct)
			protected.POST("/products", storage.CreateProduct)
			protected.PATCH("/products", storage.UpdateProduct)
			protected.DELETE("/products/:id", storage.DeleteProduct)

			protected.GET("/product_types", storage.ListProductTypes)
			protected.GET("/product_type/:id", storage.GetProductType)
			protected.POST("/product_types", storage.CreateProductType)
			protected.PATCH("/pProduct_types", storage.UpdateProductType)
			protected.DELETE("/pProduct_types/:id", storage.DeleteProductType)

			//===================================================Storage
			//==================================================Checkroom Routes
			protected.GET("/checkrooms", checkroom.ListCheckroom)
			protected.GET("/checkroom/:id", checkroom.GetCheckroom)
			protected.POST("/checkrooms", checkroom.CreateCheckroom)
			protected.PATCH("/checkrooms", checkroom.UpdateCheckroom)
			protected.DELETE("/checkrooms/:id", checkroom.DeleteCheckroom)
			//Gender
			protected.GET("/damages", checkroom.ListDamage)
			protected.GET("/checkrooms/damages/:id", checkroom.GetDamage)
			protected.POST("/checkrooms/damages", checkroom.CreateDamage)
			protected.PATCH("/checkrooms/damages", checkroom.UpdateDamage)
			protected.DELETE("/checkrooms/damages/:id", checkroom.DeleteDamage)
			//StatusCR
			protected.GET("/status", checkroom.ListStatus)
			protected.GET("/checkrooms/statuscrs/:id", checkroom.GetStatus)
			protected.POST("/checkrooms/status", checkroom.CreateStatus)
			protected.PATCH("/checkrooms/status", checkroom.UpdateStatus)
			protected.DELETE("/checkrooms/status/:id", checkroom.DeleteStatus)

			//==================================================Checkroom Routes

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
