package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateMemberlevel(c *gin.Context) {
	var memberlevel entity.Memberlevel
	if err := c.ShouldBindJSON(&memberlevel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&memberlevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": memberlevel})
}

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

// DELETE 
func DeleteMemberlevel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM memberlevels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memberlevel not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH 
func UpdateMemberlevel(c *gin.Context) {
	var memberlevel entity.Memberlevel
	if err := c.ShouldBindJSON(&memberlevel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", memberlevel.ID).First(&memberlevel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	if err := entity.DB().Save(&memberlevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberlevel})
}