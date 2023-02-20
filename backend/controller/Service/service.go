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
	var storage entity.Storage

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

	// 12: ค้นหา storage ด้วย id
	if tx := entity.DB().Where("id = ?", service.StorageID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
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
		FoodItem:    service.FoodItem,
		Drink:       drink,
		DrinkItem:   service.DrinkItem,
		Storage:     storage,
		StorageItem: service.StorageItem,
		Total:       service.Total,
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
	if err := entity.DB().Preload("Food").Preload("Drink").Preload("Storage.Product").Preload("Customer").Raw("SELECT * FROM services").Find(&services).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services})
}

// GET /services/customer/:id
func ListServicesByUID(c *gin.Context) {
	var services []entity.Service
	id := c.Param("id")
	if err := entity.DB().Preload("Food").Preload("Drink").Preload("Customer").Preload("Storage.Product").Raw("SELECT * FROM services WHERE customer_id = ?", id).Find(&services).Error; err != nil {
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
	var storage entity.Storage

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

	if tx := entity.DB().Where("id = ?", service.StorageID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
		return
	}

	patchservice := entity.Service{
		Time:        service.Time, // ข้อมูลที่รับเข้ามาจาก frontend
		Food:        food,
		FoodItem:    service.FoodItem,
		Drink:       drink,
		DrinkItem:   service.DrinkItem,
		Storage:     storage,
		StorageItem: service.StorageItem,
		Customer:    customer,
		Total:       service.Total,
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

// GET /food/item/:id
func GetFoodItem(c *gin.Context) {
	var food entity.Food
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT item FROM foods WHERE id = ?", id).First(&food); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": food})
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

// GET /drinks
func ListDrinks(c *gin.Context) {
	var drinks []entity.Drink
	if err := entity.DB().Raw("SELECT * FROM drinks").Find(&drinks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": drinks})
}

// GET /drink/item/:id
func GetDrinkItem(c *gin.Context) {
	var drink entity.Drink
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT item FROM drinks WHERE id = ?", id).First(&drink); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drink not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drink})
}

// PATCH /drinks
func UpdateDrink(c *gin.Context) {
	var drink entity.Drink

	if err := c.ShouldBindJSON(&drink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patchdrink := entity.Drink{
		Item: drink.Item, // ข้อมูลที่รับเข้ามาจาก frontend
	}

	if err := entity.DB().Where("id = ?", drink.ID).Updates(&patchdrink).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patchdrink})
}

// GET /accessories
func ListAccessories(c *gin.Context) {
	var accessories []entity.Storage
	if err := entity.DB().Preload("Product").Raw("SELECT * FROM storages").Find(&accessories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accessories})
}

// GET /accessorie/item/:id
func GetAccessoriesItem(c *gin.Context) {
	var accessories entity.Storage
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT quantity FROM storages WHERE id = ?", id).First(&accessories); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accessorie not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": accessories})
}

// PATCH /accessories
func UpdateAccessorie(c *gin.Context) {
	var accessories entity.Storage

	if err := c.ShouldBindJSON(&accessories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patchaccessories := entity.Storage{
		Quantity: accessories.Quantity, // ข้อมูลที่รับเข้ามาจาก frontend
	}

	if err := entity.DB().Where("id = ?", accessories.ID).Updates(&patchaccessories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patchaccessories})
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

// GET /food/price/:id
func GetFoodPrice(c *gin.Context) {
	var food entity.Food
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT price FROM foods WHERE id = ?", id).First(&food); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": food})
}

// GET /drink/price/:id
func GetDrinkPrice(c *gin.Context) {
	var drink entity.Drink
	id := c.Param("id")
	if tx := entity.DB().Raw("SELECT price FROM drinks WHERE id = ?", id).First(&drink); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drink not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drink})
}

// GET /accessorie/price/:id
func GetAccessoriePrice(c *gin.Context) {
	var accessorie entity.Product
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT price From products WHERE id = (SELECT product_id FROM storages WHERE id = ?)", id).Scan(&accessorie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accessorie})
}

// ("SELECT price From products WHERE id = (SELECT product_id FROM storages WHERE id = ?)
