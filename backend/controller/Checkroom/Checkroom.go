package controller

import (
	"net/http"

	"github.com/sut65/team03/entity"
	"github.com/gin-gonic/gin"
	
)

func CreateCheckroom(c *gin.Context) {
	var checkroom entity.Checkroom
	var room entity.Room
	var product entity.Product
	var damage entity.Damage
	var status entity.Status
	var employee entity.Employee

	if err := c.ShouldBindJSON(&checkroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error not access": err.Error()})
		return
	}
	//ค้นหา Room ด้วย id
	if tx := entity.DB().Where("id = ?", checkroom.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Roomzone not found"})
		return
	}
	
	//ค้นหา Product ด้วย id
	if tx := entity.DB().Where("id = ?", checkroom.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	//ค้นหา Damage ด้วย id
	if tx := entity.DB().Where("id = ?", checkroom.DamageID).First(&damage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Damage not found"})
		return
	}
	//ค้นหา Status ด้วย id
	if tx := entity.DB().Where("id = ?", checkroom.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}
	//ค้นหา Employee ด้วย id
	 if tx := entity.DB().Where("id = ?", checkroom.EmployeeID).First(&employee); tx.RowsAffected == 0 {
	 	c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
	 	return
	 }
	

	//create entity checkroom
	checkr := entity.Checkroom{
		Room:      room,
		Product: 	   product,
		Damage:   	   damage,
		Status:   	   status,
		Date:          checkroom.Date,
		Employee:       employee,
		
	}
	//save checkroom
	if err := entity.DB().Create(&checkr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkr})
}

// GET Checkroom/:id
func GetCheckroom(c *gin.Context) {
	var checkroom entity.Checkroom
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM checkrooms WHERE id=?", id).Scan(&checkroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkroom})
}

// GET checkroom
func ListCheckroom(c *gin.Context) {
	var checkrooms []entity.Checkroom
	if err := entity.DB().Preload("Room").Preload("Product").Preload("Damage").Preload("Status").Preload("Employee").Raw("SELECT * FROM checkrooms").Find(&checkrooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkrooms})
}
// DELETE /checkrooms/:id
func DeleteCheckroom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM checkrooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkroom not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /checkrooms
func UpdateCheckroom(c *gin.Context) {
	var checkroom entity.Checkroom
	if err := c.ShouldBindJSON(&checkroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", checkroom.ID).First(&checkroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	if err := entity.DB().Save(&checkroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkroom})
}
