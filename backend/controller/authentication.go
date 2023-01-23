package controller

import (
	"net/http"

	"github.com/sut65/team03/entity"
	"github.com/sut65/team03/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`//ไฟล์ json ให้อ่านข้อมูลง่ายขึ้น
}

// OfficerResponse token response
type OfficerResponse struct {
	Token    string        `json:"token"`
	ID       uint          `json:"id"`
	Officer   entity.Officer `json:"user"` //สร้างเพื่อแยกOfficer
	RoleName string        `json:"role"`
}

// EmployeeResponse token response
type EmployeeResponse struct {
	Token    string        `json:"token"`
	ID       uint          `json:"id"`
	Employee   entity.Employee `json:"user"` //สร้างเพื่อแยกEmployee
	RoleName string        `json:"role"`
}

// CustomerResponse token response
type CustomerResponse struct {
	Token    string       `json:"token"`
	ID       uint         `json:"id"`
	Customer    entity.Customer `json:"user"` //สร้างเพื่อแยกCustomer String
	RoleName string       `json:"role"`
}

// POST /login
func Login(c *gin.Context) {
	var payload LoginPayload
	var signin entity.Signin

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา employee ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if tx := entity.DB().Raw("SELECT * FROM signins WHERE username = ?", payload.Username).Preload("UserRole").
		Find(&signin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// ตรวจสอบรหัสผ่าน สิ่งที่เข้ารหัส เอามาถอดรหัส
	err := bcrypt.CompareHashAndPassword([]byte(signin.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	//เก็บ token
	signedToken, err := jwtWrapper.GenerateToken(signin.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	var OfficerRole entity.UserRole
	var EmployeeRole entity.UserRole
	var CustomerRole entity.UserRole
	entity.DB().Raw("SELECT * FROM user_roles WHERE role_name = ?", "Officer").First(&OfficerRole)
	entity.DB().Raw("SELECT * FROM user_roles WHERE role_name = ?", "Employee").First(&EmployeeRole)
	entity.DB().Raw("SELECT * FROM user_roles WHERE role_name = ?", "Customer").First(&CustomerRole)
//ตรวจสอบว่าใครเป็น หมอ หรือ ใครเป็น แอดมิน
	if signin.UserRole.RoleName == EmployeeRole.RoleName {
		var employee entity.Employee
		if tx := entity.DB().
			Raw("SELECT * FROM employees WHERE signin_id = ?", signin.ID).Find(&employee); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "employees not found"})
			return
		}

		tokenResponse := EmployeeResponse{
			Token:    signedToken,
			ID:       employee.ID,
			Employee:   employee,
			RoleName: EmployeeRole.RoleName,
		}

		c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
	}
	if signin.UserRole.RoleName == CustomerRole.RoleName {
		var customer entity.Customer
		if tx := entity.DB().
			Raw("SELECT * FROM customers WHERE signin_id = ?", signin.ID).Find(&customer); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "customers not found"})
			return
		}

		tokenResponse := CustomerResponse{
			Token:    signedToken,
			ID:       customer.ID,
			Customer:   customer,
			RoleName: CustomerRole.RoleName,
		}

		c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
	}else if signin.UserRole.RoleName == OfficerRole.RoleName {

		var officer entity.Officer

		if tx := entity.DB().
			Raw("SELECT * FROM officers WHERE signin_id = ?", signin.ID).Find(&officer); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "officers not found"})
			return
		}

		tokenResponse := OfficerResponse{
			Token:    signedToken,
			ID:       officer.ID,
			Officer:    officer,
			RoleName: OfficerRole.RoleName,
		}

		c.JSON(http.StatusOK, gin.H{"data": tokenResponse})

	}

}
