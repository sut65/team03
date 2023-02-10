package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRpRqPass(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	// ข้อมูลถูกต้องหมดทุก field
	rr := RepairReq{
		RoomID:       &test,
		RepairTypeID: &test,
		Note:         "anything",
		Time:         time.Now(),
		CustomerID:   &test,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rr)
	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())
	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestRpRqValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check note not blank", func(t *testing.T) {
		//Comment ห้ามว่าง
		rr := RepairReq{
			RoomID:       &test,
			RepairTypeID: &test,
			Note:         "",
			Time:         time.Now(),
			CustomerID:   &test,
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(rr)
		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())
		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Please enter note"))
	})

	t.Run("check time not blank", func(t *testing.T) {
		rr := RepairReq{
			RoomID:       &test,
			RepairTypeID: &test,
			Note:         "anything",
			Time:         time.Time{},
			CustomerID:   &test,
		}

		ok, err := govalidator.ValidateStruct(rr)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select time"))
	})

	t.Run("check room not blank", func(t *testing.T) {
		rr := RepairReq{
			RoomID:       nil,
			RepairTypeID: &test,
			Note:         "anything",
			Time:         time.Now(),
			CustomerID:   &test,
		}

		ok, err := govalidator.ValidateStruct(rr)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select room"))
	})

	t.Run("check type not blank", func(t *testing.T) {
		rr := RepairReq{
			RoomID:       &test,
			RepairTypeID: nil,
			Note:         "anything",
			Time:         time.Now(),
			CustomerID:   &test,
		}

		ok, err := govalidator.ValidateStruct(rr)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select type of problem"))
	})

	t.Run("check customer not blank", func(t *testing.T) {
		rr := RepairReq{
			RoomID:       &test,
			RepairTypeID: &test,
			Note:         "anything",
			Time:         time.Now(),
			CustomerID:   nil,
		}

		ok, err := govalidator.ValidateStruct(rr)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Signin"))
	})
}

func TestRpRqLength(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	test_data := ""
	for i := 0; i < 201; i++ {
		test_data += "A"
	}

	// ข้อมูลถูกต้องบาง field
	rr := RepairReq{
		RoomID:       &test,
		RepairTypeID: &test,
		Note:         test_data,
		Time:         time.Now(),
		CustomerID:   &test,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rr)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Note length must be between 0 - 200"))
}

func TestNoteNoSpecial(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	// ข้อมูลถูกต้องบาง field
	rr := RepairReq{
		RoomID:       &test,
		RepairTypeID: &test,
		Note:         "lol+++",
		Time:         time.Now(),
		CustomerID:   &test,
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(rr)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Note must not contain special characters"))

}
