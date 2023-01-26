package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateNametitle(c *gin.Context) { 
	var nametitle entity.Nametitle
	if err := c.ShouldBindJSON(&nametitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&nametitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": nametitle})
}

// GET /nametitles/:id
func GetNametitle(c *gin.Context) {
	var nametitle entity.Nametitle
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM nametitles WHERE id = ?", id).Scan(&nametitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nametitle})
}

// GET /nametitles
func ListNametitle(c *gin.Context) {
	var nametitles []entity.Nametitle

	if err := entity.DB().Raw("SELECT * FROM nametitles").First(&nametitles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nametitles})
}

// DELETE 
func DeleteNametitle(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM nametitles WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nametitle not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH 
func UpdateNametitle(c *gin.Context) {
	var nametitle entity.Nametitle
	if err := c.ShouldBindJSON(&nametitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", nametitle.ID).First(&nametitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	if err := entity.DB().Save(&nametitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nametitle})
}