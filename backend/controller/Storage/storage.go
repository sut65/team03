package controller

import (
	"github.com/sut65/team03/entity"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

//Product...................................................................
// POST /Products

func CreateProduct(c *gin.Context) {

	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /product/:id
func GetProduct(c *gin.Context) {

	var product entity.Product

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM products WHERE id = ?", id).Scan(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products
func ListProducts(c *gin.Context) {

	var products []entity.Product

	if err := entity.DB().Raw("SELECT * FROM products").Scan(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// DELETE /products/:id
func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /products
func UpdateProduct(c *gin.Context) {

	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", product.ID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	if err := entity.DB().Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

//ProductType...................................................................
// POST /ProductTypes

func CreateProductType(c *gin.Context) {

	var producttype entity.ProductType

	if err := c.ShouldBindJSON(&producttype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": producttype})
}

// GET /ProductType/:id
func GetProductType(c *gin.Context) {

	var producttype entity.ProductType

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM product_types WHERE id = ?", id).Scan(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": producttype})
}

// GET /ProductTypes
func ListProductTypes(c *gin.Context) {

	var producttypes []entity.ProductType

	if err := entity.DB().Raw("SELECT * FROM product_types").Scan(&producttypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": producttypes})
}

// DELETE /ProductTypes/:id
func DeleteProductType(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM product_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ProductTypes
func UpdateProductType(c *gin.Context) {

	var producttype entity.ProductType

	if err := c.ShouldBindJSON(&producttype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", producttype.ID).First(&producttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	if err := entity.DB().Save(&producttype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": producttype})
}

// Storage...................................................................
// POST /Storages
func SetupPasswordHash(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

func CreateStorage(c *gin.Context) {

	var employee entity.Employee
	var product entity.Product
	var producttype entity.ProductType
	var storage entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9.  ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", storage.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10. ค้นหา Product ด้วย id
	if tx := entity.DB().Where("id = ?", storage.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// 11. ค้นหา ProductType ด้วย id
	if tx := entity.DB().Where("id = ?", storage.ProductTypeID).First(&producttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	// 12. สร้าง Storage
	st := entity.Storage{
		ID:          storage.ID,
		Employee:    employee,
		Product:     product,
		ProductType: producttype,
		Quantity:    storage.Quantity,
		Time:        storage.Time,
	}

	if err := entity.DB().Create(&st).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": st})
}

// GET /Storage/:id
func GetStorage(c *gin.Context) {

	var storage entity.Storage
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM storages WHERE id = ?", id).Scan(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// GET /Storages
func ListStorages(c *gin.Context) {

	var storages []entity.Storage

	if err := entity.DB().Preload("Employee").Preload("Product").Preload("ProductType").Raw("SELECT * FROM storages").Find(&storages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": storages})
}

// DELETE /Storages/:id
func DeleteStorage(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM storages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storage not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PUT Storage
func UpdateStorage(c *gin.Context) {
	var storage entity.Storage
	var sr entity.Storage

	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา Room ด้วย ID
	if tx := entity.DB().Where("id = ?", storage.ID).First(&sr); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Storage not found"})
		return
	}

	var employee entity.Employee
	var product entity.Product
	var producttype entity.ProductType

	// 9.  ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", storage.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10. ค้นหา Product ด้วย id
	if tx := entity.DB().Where("id = ?", storage.ProductID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// 11. ค้นหา ProductType ด้วย id
	if tx := entity.DB().Where("id = ?", storage.ProductTypeID).First(&producttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "producttype not found"})
		return
	}

	//ro.Time = room.Time
	sr.Product = product
	sr.ProductType = producttype
	sr.Quantity = storage.Quantity

	if err := entity.DB().Save(&sr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sr})
}
