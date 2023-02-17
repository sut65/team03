package controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateBooking(c *gin.Context) {
	var booking entity.Booking
	var customer entity.Customer
	var room entity.Room
	var branch entity.Branch

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", booking.BranchID).First(&branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	// 10: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", booking.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 11: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", booking.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	start := time.Date(booking.Start.Year(), booking.Start.Month(), booking.Start.Day(), 0, 0, 0, 0, time.UTC)
	stop := time.Date(booking.Stop.Year(), booking.Stop.Month(), booking.Stop.Day(), 0, 0, 0, 0, time.UTC)
	//for grouping
	hashBk_No := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v_%v_%v_%v", stop.Unix(), start.Unix(), room.ID, branch.ID))))

	// Loop over each day in the date range
	for d := start; d.Before(stop.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		// Create a hash string from the room_id and dayeach
		hashTx_No := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v_%v", d.Unix(), room.ID))))
		// Create a booking for each day
		// 12: สร้าง Booking
		bk := entity.Booking{
			Booking_Number: hashBk_No,
			Tx_No:          hashTx_No,
			Branch:         branch,
			Room:           room,
			Start:          start,
			Stop:           stop,
			Customer:       customer,
			Total:          float64(room.Amount),
			DayEach:        d,
		}

		// Save the booking to the database
		//13: save
		if err := entity.DB().Create(&bk).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"data": booking})
}

// GET /booking/:id
func GetBooking(c *gin.Context) {
	id := c.Param("id")
	var bookings struct {
		TotalAmount float64
		entity.Booking
	}
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT *, SUM(total) as total_amount FROM bookings WHERE id = ? GROUP BY booking_number", id).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /bookings
func ListBookings(c *gin.Context) {
	var b_total []struct {
		TotalAmount float64
		entity.Booking
	}
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT *, SUM(total) as total_amount FROM bookings GROUP BY booking_number").Find(&b_total).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": b_total})
}

// GET /bookings/user/:id
func ListBookingsByUID(c *gin.Context) {
	id := c.Param("id")
	var bookings []struct {
		TotalAmount float64
		entity.Booking
	}

	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT *, SUM(total) as total_amount FROM bookings WHERE customer_id = ? GROUP BY booking_number", id).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /bookingbydate
func ListBookingsBydate(c *gin.Context) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	//formattedDate := today.Format("2006-01-02")
	var bookings []struct {
		TotalAmount float64
		entity.Booking
	}
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT *, SUM(total) as total_amount FROM bookings WHERE start = ? GROUP BY booking_number", today).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /bookingtotalgroupbydate
func ListBookingsTotalbyCID(c *gin.Context) {
	var b_total []struct {
		CustomerID  uint64
		TotalAmount float64
	}
	if err := entity.DB().Raw("SELECT *, customer_id, SUM(total) as total_amount FROM bookings GROUP BY customer_id").Scan(&b_total).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b_total})
}

// DELETE /bookings/:id
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bookings WHERE booking_number = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bookings
func UpdateBooking(c *gin.Context) {
	var booking entity.Booking
	var customer entity.Customer
	var room entity.Room
	var branch entity.Branch

	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	// จำเป็นต้องลบเพราะว่าให้ Hash เมื่อถูกแก้ไขค่า Hash อาจจะเปลียนได้
	if tx := entity.DB().Exec("DELETE FROM bookings WHERE booking_number = ?", booking.Booking_Number); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา branch ด้วย id
	if tx := entity.DB().Where("id = ?", booking.BranchID).First(&branch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	// 10: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", booking.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 11: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", booking.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	start := time.Date(booking.Start.Year(), booking.Start.Month(), booking.Start.Day(), 0, 0, 0, 0, time.UTC)
	stop := time.Date(booking.Stop.Year(), booking.Stop.Month(), booking.Stop.Day(), 0, 0, 0, 0, time.UTC)
	//for grouping
	hashBk_No := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v_%v_%v_%v", stop.Unix(), start.Unix(), room.ID, branch.ID))))

	// Loop over each day in the date range
	for d := start; d.Before(stop.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		// Create a hash string from the room_id and dayeach
		hashTx_No := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v_%v", d.Unix(), room.ID))))
		// Create a booking for each day
		// 12: สร้าง Booking
		bk := entity.Booking{
			Booking_Number: hashBk_No,
			Tx_No:          hashTx_No,
			Branch:         branch,
			Room:           room,
			Start:          start,
			Stop:           stop,
			Customer:       customer,
			Total:          float64(room.Amount),
			DayEach:        d,
		}

		// Save the booking to the database
		//13: save
		if err := entity.DB().Create(&bk).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"data": booking})
}
