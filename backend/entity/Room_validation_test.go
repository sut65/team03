package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRoomPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	room := Room{
		Room_No: "A10",      // format: A-D + ตัวเลข2ตัว
		Time:    time.Now(), // เป็นปัจจุบัน +- 3 นาที
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(room)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err ต้องเป็น nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestRoomRoom_No_Format(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูล Room_NO ไม่ถูกต้องตาม Format
	room := Room{
		Room_No: "12", // format is A-D ตามด้วยตัวเลข 2 ตัว
		Time:    time.Now().Add(22 * time.Hour),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(room)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("หมายเลขห้องต้องขึ้นต้นด้วย A-D ตามด้วยตัวเลข 2 หลัก"))
}

func TestRoom_No__Null(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูล Room_No ไม่ถูกต้องตาม Format
	room := Room{
		Room_No: "", // Null
		Time:    time.Now().Add(22 * time.Hour),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(room)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("กรุณากรอกหมายเลขห้อง"))
}

func TestRoomTime(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูล Time เป็น อดีต
	room := Room{
		Room_No: "A10",
		Time:    time.Date(2020, 3, 22, 0, 0, 0, 0, time.UTC), // เป็นอดีต
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(room)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("วันที่และเวลาไม่ถูกต้อง"))
}
