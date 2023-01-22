package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

// POST /payment
func CreatePayment(c *gin.Context) {
	var payment entity.Payment
	var customer entity.Customer
	var paymentmethod entity.PaymentMethod
	var place entity.Place

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา paymentmethod ด้วย id
	if tx := entity.DB().Where("id = ?", payment.PaymentMethodID).First(&paymentmethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymentmethod not found"})
		return
	}

	// 14: ค้นหา place ด้วย id
	if tx := entity.DB().Where("id = ?", payment.PlaceID).First(&place); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "place not found"})
		return
	}

	// 15: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", payment.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 16: สร้าง Payment
	pm := entity.Payment{
		Customer:      customer,
		PaymentMethod: paymentmethod,
		Place:         place,
	}

	//17: save
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": pm})
}

// GET /payment/:id
func GetPayment(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// GET /payments
func ListPayments(c *gin.Context) {
	var payments []entity.Payment
	if err := entity.DB().Preload("PaymentMethod").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments").Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// เผื่อแก้ไข ** ตรง userID
// GET /payment/user/:id
func ListPaymentByUID(c *gin.Context) {
	var payments []entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments WHERE userid = ?", id).Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}
