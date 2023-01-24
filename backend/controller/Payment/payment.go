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
	var crypto entity.Crypto
	var bank entity.Bank

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

	// 12: ค้นหา crypto ด้วย id
	if tx := entity.DB().Where("id = ?", payment.CryptoID).First(&crypto); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "crypto not found"})
		return
	}

	// 13: ค้นหา bank ด้วย id
	if tx := entity.DB().Where("id = ?", payment.BankID).First(&bank); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
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
		Crypto:        crypto,
		Bank:          bank,
		Place:         place,
		Time:          payment.Time,
		Picture:       payment.Picture,
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
	if err := entity.DB().Preload("PaymentMethod").Preload("Crypto").Preload("Bank").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments").Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// GET /payment/customer/:id
func ListPaymentByUID(c *gin.Context) {
	var payments []entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("PaymentMethod").Preload("Crypto").Preload("Bank").Preload("Place").Preload("Customer").Raw("SELECT * FROM payments WHERE customer_id = ?", id).Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}
