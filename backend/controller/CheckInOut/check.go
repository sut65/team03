package controller

import (
	"github.com/sut65/team03/entity"
	//"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

//CheckInOutStatus...................................................................
// POST /CheckInOutStatuss

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

// GET /CheckInOutStatus/:id
func GetCheckInOutStatus(c *gin.Context) {

	var CheckInOutStatus entity.CheckInOutStatus

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM CheckInOutStatuss WHERE id = ?", id).Scan(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}

// GET /CheckInOutStatuss
func ListCheckInOutStatuss(c *gin.Context) {

	var CheckInOutStatuss []entity.CheckInOutStatus

	if err := entity.DB().Raw("SELECT * FROM CheckInOutStatuss").Scan(&CheckInOutStatuss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatuss})
}

// DELETE /CheckInOutStatuss/:id
func DeleteCheckInOutStatus(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM CheckInOutStatuss WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CheckInOutStatus not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /CheckInOutStatuss
func UpdateCheckInOutStatus(c *gin.Context) {

	var CheckInOutStatus entity.CheckInOutStatus

	if err := c.ShouldBindJSON(&CheckInOutStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", CheckInOutStatus.ID).First(&CheckInOutStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CheckInOutStatus not found"})
		return
	}

	if err := entity.DB().Save(&CheckInOutStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CheckInOutStatus})
}
