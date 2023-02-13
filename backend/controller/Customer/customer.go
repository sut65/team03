package controller

import (
	"net/http"

	"github.com/sut65/team03/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/asaskevich/govalidator"
)

func SetupPasswordHash(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

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

	var userrole entity.UserRole
	if err := entity.DB().Model(&entity.UserRole{}).Where("role_name = ?", "Customer").First(&userrole).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customers role not found"})
		return
	}

	createuserlogin := entity.Signin{
		Username: customer.Email,
		Password: SetupPasswordHash(customer.Password),
		UserRole: userrole,
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
		Signin:      createuserlogin,
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
	if err := entity.DB().Preload("Nametitle").Preload("Gender").Preload("Province").Raw("SELECT * FROM customers").Find(&customers).Error; err != nil {
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
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // แทรกการ validate ไว้ช่วงนี้ของ controller
	// if _, err := govalidator.ValidateStruct(customer); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := entity.DB().Save(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}