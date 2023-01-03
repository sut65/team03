package controller

 

import (

        "github.com/sut65/team03/entity"

        "github.com/gin-gonic/gin"

        "net/http"

)
//Officer...................................................................
// POST /Officers

func CreateOfficer(c *gin.Context) {

	var officer entity.Officer

	if err := c.ShouldBindJSON(&officer); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&officer).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// GET /officer/:id
func GetOfficer(c *gin.Context) {

	var officer entity.Officer

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM officers WHERE id = ?", id).Scan(&officer).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// GET /officers
func ListOfficers(c *gin.Context) {

	var officers []entity.Officer

	if err := entity.DB().Raw("SELECT * FROM officers").Scan(&officers).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// DELETE /officers/:id
func DeleteOfficer(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM officers WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /officers
func UpdateOfficer(c *gin.Context) {

	var officer entity.Officer

	if err := c.ShouldBindJSON(&officer); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", officer.ID).First(&officer); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		   return
	}

	if err := entity.DB().Save(&officer).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": officer})
}

//Department...................................................................
// POST /Departments

func CreateDepartment(c *gin.Context) {

	var department entity.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&department).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// GET /department/:id
func GetDepartment(c *gin.Context) {

	var department entity.Department

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM departments WHERE id = ?", id).Scan(&department).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// GET /departments
func ListDepartments(c *gin.Context) {

	var departments []entity.Department

	if err := entity.DB().Raw("SELECT * FROM departments").Scan(&departments).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": departments})
}

// DELETE /departments/:id
func DeleteDepartment(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM departments WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /departments
func UpdateDepartment(c *gin.Context) {

	var department entity.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", department.ID).First(&department); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		   return
	}

	if err := entity.DB().Save(&department).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

//Position...................................................................
// POST /Positions

func CreatePosition(c *gin.Context) {

	var position entity.Position

	if err := c.ShouldBindJSON(&position); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&position).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": position})
}

// GET /Position/:id
func GetPosition(c *gin.Context) {

	var position entity.Position

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM positions WHERE id = ?", id).Scan(&position).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": position})
}

// GET /Positions
func ListPositions(c *gin.Context) {

	var positions []entity.Position

	if err := entity.DB().Raw("SELECT * FROM positions").Scan(&positions).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": positions})
}

// DELETE /Positions/:id
func DeletePosition(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM positions WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Positions
func UpdatePosition(c *gin.Context) {

	var position entity.Position

	if err := c.ShouldBindJSON(&position); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", position.ID).First(&position); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
		   return
	}

	if err := entity.DB().Save(&position).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": position})
}

//Location...................................................................
// POST /Locations

func CreateLocation(c *gin.Context) {

	var location entity.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&location).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": location})
}

// GET /Location/:id
func GetLocation(c *gin.Context) {

	var location entity.Location

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM locations WHERE id = ?", id).Scan(&location).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": location})
}

// GET /Locations
func ListLocations(c *gin.Context) {

	var locations []entity.Location

	if err := entity.DB().Raw("SELECT * FROM locations").Scan(&locations).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": locations})
}

// DELETE /Locations/:id
func DeleteLocation(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM locations WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Locations
func UpdateLocation(c *gin.Context) {

	var location entity.Location

	if err := c.ShouldBindJSON(&location); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", location.ID).First(&location); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})
		   return
	}

	if err := entity.DB().Save(&location).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": location})
}

//Employee...................................................................
// POST /Employees

func CreateEmployee(c *gin.Context) {

	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&employee).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /Employee/:id
func GetEmployee(c *gin.Context) {

	var employee entity.Employee

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).Scan(&employee).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /Employees
func ListEmployees(c *gin.Context) {

	var employees []entity.Employee

	if err := entity.DB().Raw("SELECT * FROM employees").Scan(&employees).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// DELETE /Employees/:id
func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Employees
func UpdateEmployee(c *gin.Context) {

	var employee entity.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {
		   c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		   return
	}

	if err := entity.DB().Save(&employee).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}