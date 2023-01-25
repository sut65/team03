package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateCHK_Payment(c *gin.Context) {
	var chk_payment entity.CHK_Payment
	var employee entity.Employee
	var status entity.CHK_PaymentStatus
	var payment entity.Payment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร chk_payment
	if err := c.ShouldBindJSON(&chk_payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา payment ด้วย id
	if tx := entity.DB().Where("id = ?", chk_payment.PaymentID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}

	// 10: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", chk_payment.CHK_PaymentStatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	// 11: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", chk_payment.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 12: สร้าง CHK_Payment
	chk_p := entity.CHK_Payment{
		// waiting for payment
		// Payment:   payment,
		CHK_PaymentStatus: status,
		Date_time:         chk_payment.Date_time,
		Amount:            chk_payment.Amount,
		Employee:          employee,
	}

	//13: save
	if err := entity.DB().Create(&chk_p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": chk_p})
}

// GET /chk_payment/:id
func GetCHK_Payment(c *gin.Context) {
	var chk_payment entity.CHK_Payment
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&chk_payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chk_payment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": chk_payment})
}

// GET /chk_payments
func ListCHK_Payments(c *gin.Context) {
	var chk_payments entity.CHK_Payment
	if err := entity.DB().Preload("Payment").Preload("CHK_PaymentStatus").Preload("Employee").Raw("SELECT * FROM chk_payments").Find(&chk_payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chk_payments})
}

// DELETE /chk_payments/:id
func DeleteCHK_Payment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM chk_payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chk_payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /chk_payments
func UpdateCHK_Payment(c *gin.Context) {
	var chk_payment entity.CHK_Payment
	if err := c.ShouldBindJSON(&chk_payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", chk_payment.ID).First(&chk_payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chk_payment not found"})
		return
	}

	if err := entity.DB().Save(&chk_payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chk_payment})
}
