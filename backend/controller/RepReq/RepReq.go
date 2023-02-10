package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/sut65/team03/entity"

	//"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"net/http"
)

//RepairType...................................................................
// POST /repairtype

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

// GET /repairtype/:id
func GetRepairType(c *gin.Context) {

	var RepairType entity.RepairType

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM repair_types WHERE id = ?", id).Scan(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// GET /repairtypes
func ListRepairTypes(c *gin.Context) {

	var RepairType []entity.RepairType
	if err := entity.DB().Raw("SELECT * FROM repair_types").Scan(&RepairType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RepairType})
}

// DELETE /repairtype/:id
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

// PATCH /repairtype
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

// RepairReq...................................................................
// POST /repairreq
func CreateRepairReq(c *gin.Context) {

	var rr entity.RepairReq
	var room entity.Room
	var rtype entity.RepairType
	var customer entity.Customer

	// ผลลัพธ์ที่ได้จาก ui จะถูก bind เข้าตัวแปร rr
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(rr); err != nil {
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

	// ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", rr.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	//if room id where booking

	// 12: สร้าง CheckInOut
	RepairReq := entity.RepairReq{
		Room:       room,
		RepairType: rtype,
		Note:       rr.Note,
		Time:       rr.Time,
		Customer:   customer,
	}

	// 13: save
	if err := entity.DB().Create(&RepairReq).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": RepairReq})
}

// GET /repairreq/:id
func GetRepairReqByCid(c *gin.Context) {

	var rr []entity.RepairReq
	id := c.Param("id")

	if err := entity.DB().Preload("Room").Preload("RepairType").Preload("Customer").Raw("SELECT * FROM repair_reqs WHERE customer_id = ?", id).Find(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rr})
}

// GET /repairreqs
func ListRepairReqs(c *gin.Context) {

	var rr []entity.RepairReq

	if err := entity.DB().Preload("Room").Preload("RepairType").Preload("Customer").Raw("SELECT * FROM repair_reqs").Find(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rr})
}

// GET /rooms/customer/:id
func GetListRoomByCID(c *gin.Context) {
	var booking []entity.Booking
	id := c.Param("id")
	if err := entity.DB().Preload("Branch").Preload("Room").Preload("Customer").Raw("SELECT * FROM bookings WHERE customer_id = ?", id).Find(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// DELETE /repairreq/:id
func DeleteRepairReq(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM repair_reqs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_reqs not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /repairreq
func UpdateRepairReq(c *gin.Context) {

	//main
	var rr entity.RepairReq
	var rrO entity.RepairReq
	//relation
	var room entity.Room
	var rtype entity.RepairType
	var customer entity.Customer

	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	//check rr created ?
	if tx := entity.DB().Where("id = ?", rr.ID).First(&rrO); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair_reqs id = %d not found"})
		return
	}

	//check min field
	if rr.Note != "" || rr.RoomID != nil || rr.RepairTypeID != nil {

		if rr.Note == "" {
			rr.Note = rrO.Note
		}

		//room
		if rr.RoomID != nil {
			if tx := entity.DB().Where("id = ?", rr.RoomID).First(&room); tx.RowsAffected == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
				return
			}
			rr.Room = room
		} else {
			rr.RoomID = rrO.RoomID
		}

		// customer
		if rr.CustomerID != nil {
			if tx := entity.DB().Where("id = ?", rr.CustomerID).First(&customer); tx.RowsAffected == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
				return
			}
			rr.Customer = customer
		} else {
			rr.CustomerID = rrO.CustomerID
		}

		//new type
		if rr.RepairTypeID != nil {
			if tx := entity.DB().Where("id = ?", rr.RepairTypeID).First(&rtype); tx.RowsAffected == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Type not found"})
				return
			}
			rr.RepairType = rtype
		} else {
			rr.RepairTypeID = rrO.RepairTypeID
		}

		// แทรกการ validate ไว้ช่วงนี้ของ controller
		if _, err := govalidator.ValidateStruct(rr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := entity.DB().Save(&rr).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Update Success",
			"data":   rr,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter at least 1 field"})
		return
	}
}
