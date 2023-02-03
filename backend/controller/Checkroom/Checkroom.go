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

// PATCH /checkrooms/:id
func UpdateCheckroom(c *gin.Context) {
// 	//main
// 	var checkroom entity.Checkroom
// 	var checkroomUp entity.Checkroom

// 	var room entity.Room
// 	var product entity.Product
// 	var damage entity.Damage
// 	var status entity.Status
// 	var employee entity.Employee

// 	if err := c.ShouldBindJSON(&checkroom); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	//check created ?
// 	if tx := entity.DB().Where("id = ?", checkroom.ID).First(&checkroom); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("checkrooms id = %d not found", checkroom.ID)})
// 		return
// 	}

// 	if checkroom.Date.String() == "0001-01-01 00:00:00 +0000 UTC" {
// 		checkroom.Date = checkroomUp.Date
// 	}


// 	//room
// 	if checkroom.RoomID != nil {
// 		if tx := entity.DB().Where("id = ?", checkroom.RoomID).First(&room); tx.RowsAffected == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
// 			return
// 		}
// 		checkroom.Room = room
// 	} else {
// 		checkroom.RoomID = checkroomUp.RoomID
// 	}

// 	//product
// 	if checkroom.ProductID != nil {
// 		if tx := entity.DB().Where("id = ?", checkroom.ProductID).First(&product); tx.RowsAffected == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
// 			return
// 		}
// 		checkroom.Product = product
// 	} else {
// 		checkroom.ProductID = checkroomUp.ProductID
// 	}

// 	//damage
// 	if checkroom.DamageID != nil {
// 		if tx := entity.DB().Where("id = ?", checkroom.DamageID).First(&damage); tx.RowsAffected == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Damage not found"})
// 			return
// 		}
// 		checkroom.Damage = damage
// 	} else {
// 		checkroom.DamageID = checkroomUp.DamageID
// 	}

// 	//status
// 	if checkroom.StatusID != nil {
// 		if tx := entity.DB().Where("id = ?", checkroom.StatusID).First(&status); tx.RowsAffected == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
// 			return
// 		}
// 		checkroom.Status = status
// 	} else {
// 		checkroom.StatusID = checkroomUp.StatusID
// 	}


// 	// employee
// 	if checkroom.EmployeeID != nil {
// 		if tx := entity.DB().Where("id = ?", checkroom.EmployeeID).First(&employee); tx.RowsAffected == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
// 			return
// 		}
// 		checkroom.Employee = employee
// 	} else {
// 		checkroom.EmployeeID = checkroomUp.EmployeeID
// 	}



// 	if err := entity.DB().Save(&checkroom).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "Update Success",
// 		"data":   checkroom,
// 	})
var checkroom entity.Checkroom
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&checkroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkroom not found"})
		return
	}

	if err := c.ShouldBindJSON(&checkroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Save(&checkroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkroom})

}

