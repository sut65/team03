package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStoragePass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	storage := Storage{
		Quantity: 50,         // format: จำนวนเต็มบวก
		Time:     time.Now(), // เป็นปัจจุบัน +- 3 นาที
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(storage)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err ต้องเป็น nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestQuantity_Zero(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูล Room_No ไม่ถูกต้องตาม Format
	storage := Storage{
		Quantity: 0, // เป็นศูนย์
		Time:     time.Now().Add(22 * time.Hour),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(storage)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("กรุณากรอกจำนวนที่มากกว่าศูนย์"))
}

func TesQuantity_CannotBeNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Quantity cannot be negative number", func(t *testing.T) {
		storage := Storage{
			Quantity: -50, // เป็นจำนวนติดลบ
			Time:     time.Now().Add(22 * time.Hour),
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(storage)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())
		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("กรุณากรอกจำนวนเป็นจำนวนเต็มบวก"))
	})
}

func TestStorageTime(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูล Time เป็น อดีต
	storage := Storage{
		Quantity: 50,
		Time:     time.Date(2020, 3, 22, 0, 0, 0, 0, time.UTC), // เป็นอดีต
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(storage)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("วันที่และเวลาไม่ถูกต้อง"))
}
