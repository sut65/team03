package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

// POST /chk_payment/statuses
func CreateStatus(c *gin.Context) {
	var chk_paymentstatus entity.CHK_PaymentStatus
	if err := c.ShouldBindJSON(&chk_paymentstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&chk_paymentstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": chk_paymentstatus})
}

// GET /chk_payment/status/:id
func GetStatus(c *gin.Context) {
	var chk_paymentstatus entity.CHK_PaymentStatus
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&chk_paymentstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chk_paymentstatus})
}

// GET /chk_payment/statuses
func ListStatuses(c *gin.Context) {
	var chk_paymentstatuses []entity.CHK_PaymentStatus
	if err := entity.DB().Raw("SELECT * FROM chk_payment_statuses").Scan(&chk_paymentstatuses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chk_paymentstatuses})
}

// DELETE /chk_payment/statuses/:id
func DeleteStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM chk_payment_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /chk_payment/statuses
func UpdateStatus(c *gin.Context) {
	var chk_paymentstatus entity.CHK_PaymentStatus
	if err := c.ShouldBindJSON(&chk_paymentstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", chk_paymentstatus.ID).First(&chk_paymentstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	if err := entity.DB().Save(&chk_paymentstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": chk_paymentstatus})
}
