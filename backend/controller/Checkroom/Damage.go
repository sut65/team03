package controller
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

func CreateDamage(c *gin.Context) {
	var damage entity.Damage
	if err := c.ShouldBindJSON(&damage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&damage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": damage})
}

// GET /damage/:id
func GetDamage(c *gin.Context) {
	var damage entity.Damage
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM damages WHERE id = ?", id).Scan(&damage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": damage})
}

// GET /damages
func ListDamage(c *gin.Context) {
	var damage []entity.Damage

	if err := entity.DB().Raw("SELECT * FROM damages").Scan(&damage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": damage})
}

// DELETE /damages/:id
func DeleteDamage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM damages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "damage not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /damages
func UpdateDamage(c *gin.Context) {
	var damage entity.Damage
	if err := c.ShouldBindJSON(&damage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", damage.ID).First(&damage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "damage not found"})
		return
	}

	if err := entity.DB().Save(&damage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": damage})
}

