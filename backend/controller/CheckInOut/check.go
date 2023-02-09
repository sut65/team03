package controller

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
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

	var cio entity.CheckInOut //main
	//var co entity.CheckInOut //new
	var booking entity.Booking
	var status entity.CheckInOutStatus
	var emp entity.Employee

	// ผลลัพธ์ที่ได้จาก ui จะถูก bind เข้าตัวแปร cio
	if err := c.ShouldBindJSON(&cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(cio); err != nil {
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

	//fmt.Println(cio.CheckOutTime)

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

	//if err := entity.DB().Raw("SELECT * FROM Check_In_Outs").Scan(&CheckInOut).Error; err != nil {
	if err := entity.DB().Preload("Booking").Preload("CheckInOutStatus").Preload("Employee").Raw("SELECT * FROM Check_In_Outs").Find(&CheckInOut).Error; err != nil {
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

// PATCH /checkin
func UpdateCheckIn(c *gin.Context) {
	//main
	var cio entity.CheckInOut
	var cioOld entity.CheckInOut

	//relation
	//var booking entity.Booking
	//var status entity.CheckInOutStatus
	var emp entity.Employee

	if err := c.ShouldBindJSON(&cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check cio haved?
	if tx := entity.DB().Where("id = ?", cio.ID).First(&cioOld); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("CIO id = %d not found", cio.ID)})
		return
	}

	cio.CheckInOutStatusID = cioOld.CheckInOutStatusID

	if cio.CheckInTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
		cio.CheckInTime = cioOld.CheckInTime
	}

	if cio.CheckOutTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
		cio.CheckOutTime = cioOld.CheckOutTime
	}

	if cio.BookingID == nil {
		cio.BookingID = cioOld.BookingID
	}

	//new emp
	if tx := entity.DB().Where("id = ?", cio.EmployeeID).First(&emp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	cio.Employee = emp

	if err := entity.DB().Save(&cio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "Update Success",
		"data":        cio,
		"cioOld data": cioOld,
	})
}

// PATCH /checkout
func UpdateCheckOut(c *gin.Context) {
	//main
	var cio entity.CheckInOut
	var cioOld entity.CheckInOut

	//relation
	//var booking entity.Booking
	var status entity.CheckInOutStatus
	var emp entity.Employee

	if err := c.ShouldBindJSON(&cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check cio haved?
	if tx := entity.DB().Where("id = ?", cio.ID).First(&cioOld); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("CIO id = %d not found", cio.ID)})
		return
	}

	if tx := entity.DB().Where("id = ?", cio.CheckInOutStatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}
	cio.CheckInOutStatus = status

	if cio.CheckInTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
		cio.CheckInTime = cioOld.CheckInTime
	}

	if cio.CheckOutTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
		cio.CheckOutTime = cioOld.CheckOutTime
	}
	if cio.BookingID == nil {
		cio.BookingID = cioOld.BookingID
	}

	//new emp
	if tx := entity.DB().Where("id = ?", cio.EmployeeID).First(&emp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	cio.Employee = emp

	if err := entity.DB().Save(&cio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   cio,
	})
}

// PATCH /checkinout
func UpdateCheckInOut(c *gin.Context) {
	//main
	var cio entity.CheckInOut
	var cioOld entity.CheckInOut

	//relation
	//booking shouldn't change cuz 1-1 w/ cio
	//var booking entity.Booking
	var status entity.CheckInOutStatus
	var emp entity.Employee

	if err := c.ShouldBindJSON(&cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check cio haved?
	if tx := entity.DB().Where("id = ?", cio.ID).First(&cioOld); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("CIO id = %d not found", cio.ID)})
		return
	}

	if tx := entity.DB().Where("id = ?", cio.CheckInOutStatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	} else {
		cio.CheckInOutStatus = status
	}
	if *cioOld.CheckInOutStatusID == 2 {
		cio.CheckInOutStatusID = cioOld.CheckInOutStatusID
	}
	if status.ID == 1 {
		cio.CheckOutTime = cioOld.CheckOutTime
	} else if status.ID == 2 {
		cio.CheckInTime = cioOld.CheckInTime
	}

	if cio.BookingID == nil {
		cio.BookingID = cioOld.BookingID
	}

	//new emp
	if tx := entity.DB().Where("id = ?", cio.EmployeeID).First(&emp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	cio.Employee = emp

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Save(&cio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   cio,
	})
}

func CheckOut(c *gin.Context) {
	id := c.Param("id")
	var cio entity.CheckInOut

	//check ว่ามีตารางหมายเลขนี้มั้ย?
	if tx := entity.DB().Where("id = ?", id).First(&cio); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("CIO id = %s not found", id)})
		c.Abort()
		return
	}

	// if cio.CheckOutTime.String() == "0001-01-01 00:00:00 +0000 UTC" {
	// 	cio.CheckOutTime = time.Now()
	// }

	cio.CheckOutTime = time.Now()
	cio.CheckInOutStatus.ID = 2

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(cio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Save(&cio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   cio,
	})
}
