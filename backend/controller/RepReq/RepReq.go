package controller

import (
	"github.com/sut65/team03/entity"
	//"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

//RepairType...................................................................
// POST /RepairType

func CreateRepairType(c *gin.Context) {

	var RepairType entity.RepairType
	if err := c.ShouldBindJSON(&RepairType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// GET /RepairType/:id
func GetRepairType(c *gin.Context) {

	var RepairType entity.RepairType

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM repair_types WHERE id = ?", id).Scan(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// GET /RepairTypes
func ListRepairTypes(c *gin.Context) {

	var RepairType []entity.RepairType
	if err := entity.DB().Raw("SELECT * FROM repair_types").Scan(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// DELETE /RepairType/:id
func DeleteRepairType(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM repair_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_types not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

	// if err := entity.DB().Raw("DELETE FROM repair_types WHERE id = ?", id).Scan(&RepairType).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// PATCH /RepairType
func UpdateRepairType(c *gin.Context) {

	var RepairType entity.RepairType

	if err := c.ShouldBindJSON(&RepairType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", RepairType.ID).First(&RepairType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_types not found"})
		return
	}

	if err := entity.DB().Save(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

//RepairStatus...................................................................
// POST /RepairStatus

func CreateRepairStatus(c *gin.Context) {

	var RepairStatus entity.RepairStatus
	if err := c.ShouldBindJSON(&RepairStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&RepairStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairStatus})
}

// GET /RepairStatus/:id
func GetRepairStatus(c *gin.Context) {

	var RepairStatus entity.RepairStatus

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM repair_statuses WHERE id = ?", id).Scan(&RepairStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RepairStatus})
}

// GET /RepairStatuses
func ListRepairStatuses(c *gin.Context) {

	var RepairStatus []entity.RepairStatus
	if err := entity.DB().Raw("SELECT * FROM repair_statuses").Scan(&RepairStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairStatus})
}

// DELETE /RepairStatus/:id
func DeleteRepairStatus(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM repair_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_statuses not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

	// if err := entity.DB().Raw("DELETE FROM repair_statuses WHERE id = ?", id).Scan(&RepairStatus).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": RepairStatus})
}

// PATCH /RepairStatus
func UpdateRepairStatus(c *gin.Context) {

	var RepairStatus entity.RepairStatus

	if err := c.ShouldBindJSON(&RepairStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", RepairStatus.ID).First(&RepairStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_statuses not found"})
		return
	}

	if err := entity.DB().Save(&RepairStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairStatus})
}

// RepairReq...................................................................
// POST /RepairReq
func CreateRepairReq(c *gin.Context) {

	var rr entity.RepairReq
	var room entity.Room
	var rtype entity.RepairType
	var status entity.RepairStatus
	var customer entity.Customer

	// ผลลัพธ์ที่ได้จาก ui จะถูก bind เข้าตัวแปร rr
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", rr.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
		return
	}

	// ค้นหา type ด้วย id
	if tx := entity.DB().Where("id = ?", rr.RepairTypeID).First(&rtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type not found"})
		return
	}

	// ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", rr.RepairStatus).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	// ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", rr.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	// 12: สร้าง CheckInOut
	RepairReq := entity.RepairReq{
		Room:         room,
		RepairType:   rtype,
		Note:         rr.Note,
		Time:         rr.Time,
		RepairStatus: status,
		Customer:     customer,
	}

	// 13: save
	if err := entity.DB().Create(&RepairReq).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": RepairReq})
}

// GET /RepairReq/:id
func GetRepairReq(c *gin.Context) {

	var rr entity.RepairReq

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM repair_reqs WHERE id = ?", id).Scan(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rr})
}

// GET /RepairReqs
func ListRepairReqs(c *gin.Context) {

	var rr []entity.RepairReq

	if err := entity.DB().Raw("SELECT * FROM repair_reqs").Scan(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rr})
}

// DELETE /RepairReq/:id
func DeleteRepairReq(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM repair_reqs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_reqs not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /RepairReq
func UpdateRepairReq(c *gin.Context) {

	var rr entity.RepairReq

	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", rr.ID).First(&rr); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_reqs not found"})
		return
	}

	if err := entity.DB().Save(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rr})
}
