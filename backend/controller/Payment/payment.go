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
	var method entity.Method

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา paymentmethod ด้วย id
	if tx := entity.DB().Where("id = ?", payment.PaymentMethodID).First(&paymentmethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paymentmethod not found"})
		return
	}

	// 11: ค้นหา method ด้วย id
	if tx := entity.DB().Where("id = ?", payment.MethodID).First(&method); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "method not found"})
		return
	}

	// 12: ค้นหา place ด้วย id
	if tx := entity.DB().Where("id = ?", payment.PlaceID).First(&place); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "place not found"})
		return
	}

	// 13: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", payment.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 14: สร้าง Payment
	pm := entity.Payment{
		Customer:      customer,
		PaymentMethod: paymentmethod,
		Method:        method,
		Place:         place,
		Time:          payment.Time,
		Picture:       payment.Picture,
	}

	//15: save
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
	if err := entity.DB().Preload("PaymentMethod").Preload("Method").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments").Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// GET /payment/customer/:id
func ListPaymentByUID(c *gin.Context) {
	var payments []entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Preload("Method").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments WHERE customer_id = ?", id).Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// GET /paymentmethods
func ListPaymentMethods(c *gin.Context) {
	var paymentmethods []entity.PaymentMethod
	if err := entity.DB().Raw("SELECT * FROM payment_methods").Find(&paymentmethods).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethods})
}

// GET /methods/paymet/:id
func ListMethodsByPID(c *gin.Context) {
	var methods []entity.Method
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM methods WHERE payment_method_id = ?", id).Find(&methods).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": methods})
}

// GET /method/:id
func GetMethod(c *gin.Context) {
	var methods entity.Method
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Raw("SELECT * FROM methods WHERE id = ?", id).Find(&methods).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": methods})
}

// GET /places
func ListPlaces(c *gin.Context) {
	var places []entity.Place
	if err := entity.DB().Raw("SELECT * FROM places").Find(&places).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": places})
}
