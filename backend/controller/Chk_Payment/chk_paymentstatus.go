package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

// POST /chk_payment/statuses
func CreateStatus(c *gin.Context) {
	var status entity.CHK_PaymentStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": status})
}

// GET /chk_payment/status/:id
func GetStatus(c *gin.Context) {
	var status entity.CHK_PaymentStatus
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET /chk_payment/statuses
func ListStatuses(c *gin.Context) {
	var statuses []entity.CHK_PaymentStatus
	if err := entity.DB().Raw("SELECT * FROM chk_paymentstatuses").Scan(&statuses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statuses})
}

// DELETE /chk_payment/statuses/:id
func DeleteStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM chk_paymentstatuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /chk_payment/statuses
func UpdateStatus(c *gin.Context) {
	var status entity.CHK_PaymentStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", status.ID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check payment status not found"})
		return
	}

	if err := entity.DB().Save(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}
