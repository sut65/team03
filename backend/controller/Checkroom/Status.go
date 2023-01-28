package controller
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateStatusCR(c *gin.Context) {
	var status entity.StatusCR
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

// GET /status/:id
func GetStatusCR(c *gin.Context) {
	var status entity.StatusCR
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM statuscrs WHERE id = ?", id).Scan(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET /statuss
func ListStatusCR(c *gin.Context) {
	var status []entity.StatusCR

	if err := entity.DB().Raw("SELECT * FROM statuscrs").Scan(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// DELETE /statuss/:id
func DeleteStatusCR(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM statuscrs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /statuss
func UpdateStatusCR(c *gin.Context) {
	var status entity.StatusCR
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", status.ID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	if err := entity.DB().Save(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

