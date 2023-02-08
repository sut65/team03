package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut65/team03/entity"
)

// POST /service
func CreateService(c *gin.Context) {
	var service entity.Service
	var customer entity.Customer
	var food entity.Food
	var drink entity.Drink
	var accessorie entity.Accessories

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา food ด้วย id
	if tx := entity.DB().Where("id = ?", service.FoodID).First(&food); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food not found"})
		return
	}

	// 11: ค้นหา drink ด้วย id
	if tx := entity.DB().Where("id = ?", service.DrinkID).First(&drink); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drink not found"})
		return
	}

	// 12: ค้นหา accessories ด้วย id
	if tx := entity.DB().Where("id = ?", service.AccessoriesID).First(&accessorie); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accessories not found"})
		return
	}

	// 13: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", service.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 14: สร้าง Service
	sv := entity.Service{
		Customer:    customer,
		Time:        service.Time, // ข้อมูลที่รับเข้ามาจาก frontend
		Food:        food,
		Drink:       drink,
		Accessories: accessorie,
	}

	//15: save
	if err := entity.DB().Create(&sv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": sv})
}

// GET /service/:id
func GetService(c *gin.Context) {
	var service entity.Service
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&service); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "service not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": service})
}

// GET /services
func ListServices(c *gin.Context) {
	var services []entity.Service
	if err := entity.DB().Preload("Food").Preload("Drink").Preload("Accessories").Preload("Customer").Raw("SELECT * FROM services").Find(&services).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services})
}

// GET /services/customer/:id
func ListServicesByUID(c *gin.Context) {
	var services []entity.Service
	id := c.Param("id")
	if err := entity.DB().Preload("Food").Preload("Drink").Preload("Accessories").Preload("Customer").Raw("SELECT * FROM services WHERE customer_id = ?", id).Find(&services).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services})
}

// DELETE /services/:id
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM services WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please type some bill."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /services
func UpdateService(c *gin.Context) {
	var service entity.Service
	var customer entity.Customer
	var food entity.Food
	var drink entity.Drink
	var accessorie entity.Accessories

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", service.FoodID).First(&food); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", service.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", service.DrinkID).First(&drink); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drink not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", service.AccessoriesID).First(&accessorie); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accessories not found"})
		return
	}

	patchservice := entity.Service{
		Time:        service.Time, // ข้อมูลที่รับเข้ามาจาก frontend
		Food:        food,
		Drink:       drink,
		Accessories: accessorie,
		Customer:    customer,
	}

	if err := entity.DB().Where("id = ?", service.ID).Updates(&patchservice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})
}

// GET /foods
func ListFoods(c *gin.Context) {
	var foods []entity.Food
	if err := entity.DB().Raw("SELECT * FROM foods").Find(&foods).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foods})
}

// GET /drinks
func ListDrinks(c *gin.Context) {
	var drinks []entity.Drink
	if err := entity.DB().Raw("SELECT * FROM drinks").Find(&drinks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": drinks})
}

// GET /accessories
func ListAccessories(c *gin.Context) {
	var accessories []entity.Accessories
	if err := entity.DB().Raw("SELECT * FROM accessories").Find(&accessories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accessories})
}

// GET /room/customer/:id
func GetRoomByCID(c *gin.Context) {
	var booking entity.Booking
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT room_id FROM bookings WHERE customer_id = ?", id).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// PATCH /foods
func UpdateFood(c *gin.Context) {
	var food entity.Food

	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patchfood := entity.Food{
		Item: food.Item, // ข้อมูลที่รับเข้ามาจาก frontend
	}

	if err := entity.DB().Where("id = ?", food.ID).Updates(&patchfood).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patchfood})
}
