package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

// GET /memberlevel/:id
func GetMemberlevel(c *gin.Context) {
	var memberlevel entity.Memberlevel
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM memberlevels WHERE id = ?", id).Scan(&memberlevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": memberlevel})
}

// GET /memberlevel
func ListMemberlevel(c *gin.Context) {
	var memberlevels []entity.Memberlevel

	if err := entity.DB().Raw("SELECT * FROM memberlevels").First(&memberlevels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": memberlevels})
}

//ยังไม่เสดดด