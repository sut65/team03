package controller

import (
	"net/http"

	"github.com/sut65/team03/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/asaskevich/govalidator"
)

func CreateCustomer(c *gin.Context) {
	var gender entity.Gender
	var province entity.Province
	var nametitle entity.Nametitle
	var customer entity.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error not access": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา Gender ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gender not found"})
		return
	}
	//ค้นหา Province ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Province_ID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Province not found"})
		return
	}
	//ค้นหา Nametitle ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Nametitle_ID).First(&nametitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nametitle not found"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 14)

	//create entity customer
	cus := entity.Customer{
		Gender:      gender,
		Province:    province,
		Nametitle:   nametitle,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Password:    string(password),
		Age:         customer.Age,
		Phone:       customer.Phone,
		Email:       customer.Email,
	}
	//save customer
	if err := entity.DB().Create(&cus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cus})
}

// GET Customer/:id
func GetCustomerByID(c *gin.Context) {
	var customer entity.Customer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM customers WHERE id=?", id).Scan(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// GET All/Customers
func ListCustomers(c *gin.Context) {
	var customers []entity.Customer
	if err := entity.DB().Raw("SELECT * FROM customers").Scan(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers})
}
// DELETE /customers/:id
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM customers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /customers
func UpdateCustomer(c *gin.Context) {
	var customer entity.Customer
	var nametitle entity.Nametitle
	var gender entity.Gender
	var province entity.Province

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา nametitle ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Nametitle_ID).First(&nametitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nametitle not found"})
		return
	}

	// ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Province_ID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 14)

	// สร้าง Customer
	update := entity.Customer{
		Nametitle: nametitle,
		Gender:      gender,
		Province:     province,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Password:    string(password),
		Age:         customer.Age,
		Phone:       customer.Phone,
		Email:       customer.Email,
	}

	// บันทึก
	if err := entity.DB().Where("id = ?", customer.ID).Updates(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": update})
}
