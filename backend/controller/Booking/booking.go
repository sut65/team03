package controller

import (
	"net/http"

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

	// 12: สร้าง Booking
	bk := entity.Booking{
		Branch:   branch,
		Room:     room,
		Start:    booking.Start,
		Stop:     booking.Stop,
		Customer: customer,
	}

	//13: save
	if err := entity.DB().Create(&bk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": bk})
}

// GET /booking/:id
func GetBooking(c *gin.Context) {
	var booking entity.Booking
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// GET /bookings
func ListBookings(c *gin.Context) {
	var bookings []entity.Booking
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT * FROM bookings").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GET /bookings/user/:id
func ListBookingsByUID(c *gin.Context) {
	var bookings []entity.Booking
	id := c.Param("id")
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT * FROM bookings WHERE userid = ?", id).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// DELETE /bookings/:id
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bookings WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bookings
func UpdateBooking(c *gin.Context) {
	var booking entity.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", booking.ID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	if err := entity.DB().Save(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}
