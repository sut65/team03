package controller

import (
	"github.com/sut65/team03/entity"
	//"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

//CheckInOutStatus...................................................................
// POST /checkinoutstatus

func CreateCheckInOutStatus(c *gin.Context) {

	var CheckInOutStatus entity.CheckInOutStatus
	if err := c.ShouldBindJSON(&CheckInOutStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

// GET //checkinoutstatus/:id
func GetCheckInOutStatus(c *gin.Context) {

	var CheckInOutStatus entity.CheckInOutStatus

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Check_In_Out_Statuses WHERE id = ?", id).Scan(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

// GET /checkinoutstatuses
func ListCheckInOutStatuses(c *gin.Context) {

	var CheckInOutStatus []entity.CheckInOutStatus
	if err := entity.DB().Raw("SELECT * FROM Check_In_Out_Statuses").Scan(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

// DELETE /checkinoutstatus/:id
func DeleteCheckInOutStatus(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Check_In_Out_Statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check_In_Out_Status not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

	// if err := entity.DB().Raw("DELETE FROM Check_In_Out_Statuses WHERE id = ?", id).Scan(&CheckInOutStatus).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

// PATCH //checkinoutstatus
func UpdateCheckInOutStatus(c *gin.Context) {

	var CheckInOutStatus entity.CheckInOutStatus

	if err := c.ShouldBindJSON(&CheckInOutStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", CheckInOutStatus.ID).First(&CheckInOutStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check_In_Out_Statuses not found"})
		return
	}

	if err := entity.DB().Save(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

//CheckInOut...................................................................
// POST /checkinout

func CreateCheckInOut(c *gin.Context) {

	var cio entity.CheckInOut
	var booking entity.Booking
	var status entity.CheckInOutStatus
	var emp entity.Employee

	// ผลลัพธ์ที่ได้จาก ui จะถูก bind เข้าตัวแปร cio
	if err := c.ShouldBindJSON(&cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา booking ด้วย id
	if tx := entity.DB().Where("id = ?", cio.BookingID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found"})
		return
	}

	// ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", cio.CheckInOutStatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	// ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", cio.EmployeeID).First(&emp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}

	// 12: สร้าง CheckInOut
	CheckInOut := entity.CheckInOut{
		Booking:          booking,
		CheckInTime:      cio.CheckInTime,
		CheckOutTime:     cio.CheckOutTime,
		CheckInOutStatus: status,
		Employee:         emp,
	}

	// 13: save
	if err := entity.DB().Create(&CheckInOut).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": CheckInOut})
}

// GET /checkinout/:id
func GetCheckInOut(c *gin.Context) {

	var CheckInOut entity.CheckInOut

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Check_In_Outs WHERE id = ?", id).Scan(&CheckInOut).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOut})
}

// GET /checkinouts
func ListCheckInOuts(c *gin.Context) {

	var CheckInOut []entity.CheckInOut

	if err := entity.DB().Raw("SELECT * FROM Check_In_Outs").Scan(&CheckInOut).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOut})
}

// DELETE /checkinout/id
func DeleteCheckInOut(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Check_In_Outs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check_In_Outs not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /checkinout
func UpdateCheckInOut(c *gin.Context) {

	var CheckInOut entity.CheckInOut

	if err := c.ShouldBindJSON(&CheckInOut); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", CheckInOut.ID).First(&CheckInOut); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check_In_Outs not found"})
		return
	}

	if err := entity.DB().Save(&CheckInOut).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOut})
}
