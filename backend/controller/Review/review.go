package controller

import (
	"github.com/sut65/team03/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateReview(c *gin.Context) {

	var review entity.Review
	var customer entity.Customer
	var systemwork entity.Systemwork
	var department entity.Department



	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", review.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", review.SystemworkID).First(&systemwork); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Systemwork not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", review.DepartmentID).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Department not found"})
		return
	}

			// 12: สร้าง Review
			rv := entity.Review{
				Customer:  customer,            		 // โยงความสัมพันธ์กับ Entity customer
				Systemwork: systemwork,                  // โยงความสัมพันธ์กับ systemwork
				Department:    department,         // โยงความสัมพันธ์กับ Entity department
				Comment: review.Comment,
				Star: review.Star,
				Reviewdate: review.Reviewdate,
				Reviewimega: review.Reviewimega,
			}

	if err := entity.DB().Create(&rv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rv})
}

// GET /Review/:id
func GetReview(c *gin.Context) {

	var review entity.Review

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM reviews WHERE id = ?", id).Scan(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GET /reviews
func ListReviews(c *gin.Context) {

	var review []entity.Review

	if err := entity.DB().Raw("SELECT * FROM reviews").Scan(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DELETE /reviews/:id
func DeleteReview(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM reviews WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /reviews
func UpdateReview(c *gin.Context) {

	var review entity.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", review.ID).First(&review); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "review not found"})
		return
	}

	if err := entity.DB().Save(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

//Systemwork...................................................................
// POST /Systemworks

func CreateSystemwork(c *gin.Context) {

	var systemwork entity.Systemwork

	if err := c.ShouldBindJSON(&systemwork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&systemwork).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": systemwork})
}

// GET /systemwork/:id
func GetSystemwork(c *gin.Context) {

	var systemwork entity.Systemwork

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM systemworks WHERE id = ?", id).Scan(&systemwork).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": systemwork})
}

// GET /systemworks
func ListSystemworks(c *gin.Context) {

	var systemwork []entity.Systemwork

	if err := entity.DB().Raw("SELECT * FROM systemworks").Scan(&systemwork).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": systemwork})
}

// DELETE /systemworks/:id
func DeleteSystemwork(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM systemworks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "systemwork not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /systemworks
func UpdateSystemwork(c *gin.Context) {

	var systemwork entity.Systemwork

	if err := c.ShouldBindJSON(&systemwork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", systemwork.ID).First(&systemwork); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "systemwork not found"})
		return
	}

	if err := entity.DB().Save(&systemwork).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": systemwork})
}