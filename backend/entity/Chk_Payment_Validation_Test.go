package entity

// import (
// 	"testing"
// 	"time"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// // ตรวจสอบค่าว่างของการจองห้องพักแล้วถูกทั้งหมด
// func TestBookingNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	t.Run("check booking not blank", func(t *testing.T) {
// 		chkp := CHK_Payment{
// 			Amount:      1,
// 			Description: "",
// 		}

// 		// ตรวจสอบด้วย govalidator
// 		ok, err := govalidator.ValidateStruct(chkp)

// 		// ok ต้องเป็นค่า true แปลว่าต้องจับ error ไม่ได้
// 		g.Expect(ok).To(BeTrue())

// 		// err ต้องเป็นค่า nil แปลว่าต้องจับ error ไม่ได้
// 		g.Expect(err).To(BeNil())

// 		g.Expect(err.Error()).To(Equal("กรุณาระบุเวลาเริ่มเข้าพัก"))
// 	})

// 	t.Run("check start not blank", func(t *testing.T) {
// 		booking := Booking{
// 			BranchID: new(uint),
// 			RoomID:   new(uint),
// 			// Start:      time.Time{},
// 			Stop:       time.Time{},
// 			CustomerID: new(uint),
// 		}
// 		// ตรวจสอบด้วย govalidator
// 		ok, err := govalidator.ValidateStruct(booking)

// 		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 		g.Expect(ok).ToNot(BeTrue())

// 		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 		g.Expect(err).ToNot(BeNil())

// 		// err.Error ต้องมี error message แสดงออกมา
// 		g.Expect(err.Error()).To(Equal("กรุณาระบุเวลาเริ่มเข้าพัก"))
// 	})
// }
