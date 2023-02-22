package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

func TestCIOPass(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	// ข้อมูลถูกต้องหมดทุก field
	cio := CheckInOut{
		BookingID:          &test,
		CheckInTime:        time.Now(),
		CheckOutTime:       time.Now().AddDate(0, 0, +1),
		CheckInOutStatusID: &test,
		EmployeeID:         &test,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(cio)
	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())
	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestCheckInOutValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check check-in time not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Time{},
			CheckOutTime:       time.Now(),
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select check-in time"))
	})

	// t.Run("check check-out time not blank", func(t *testing.T) {
	// 	cio := CheckInOut{
	// 		BookingID:          &test,
	// 		CheckInTime:        time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
	// 		CheckOutTime:       time.Time{},
	// 		CheckInOutStatusID: &test,
	// 		EmployeeID:         &test,
	// 	}

	// 	ok, err := govalidator.ValidateStruct(cio)

	// 	g.Expect(ok).To(BeFalse())
	// 	g.Expect(err).To(HaveOccurred())
	// 	g.Expect(err.Error()).To(Equal("Please select check-out time"))
	// })

	t.Run("check booking not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          nil,
			CheckInTime:        time.Now(),
			CheckOutTime:       time.Now().AddDate(0, 0, +1),
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select booking"))
	})

	t.Run("check status not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Now(),
			CheckOutTime:       time.Now().AddDate(0, 0, +1),
			CheckInOutStatusID: nil,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select status"))
	})

	t.Run("check employee not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Now(),
			CheckOutTime:       time.Now().AddDate(0, 0, +1),
			CheckInOutStatusID: &test,
			EmployeeID:         nil,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Signin"))
	})
}

// ============================================================= Date Now ->
func TestValidateCheckInTimeMustBeNowNotPast(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check in time must be now not past", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Now().AddDate(0, 0, -1), //ผิด เป็นอดีต
			CheckOutTime:       time.Now(),
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid check-in time (do not be in the past)"))
	})
}

// ============================================================= Date Now ->
func TestValidateCheckInTimeMustBeNowNotFuture(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check in time must be now not future", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Now().AddDate(0, 0, 1), //ผิด เป็นอดีต
			CheckOutTime:       time.Now().AddDate(0, 0, 2),
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid check-in time (do not be in the future)"))
	})
}

// ============================================================= Stop = Start+1 ->
func TestValidateCheckOutTimeMustbeAfterStartOneMinute(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check out time must be after in time", func(t *testing.T) {

		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Now(),
			CheckOutTime:       time.Now().Add(time.Second * -1), //ผิด
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("check out time must be after in check in time"))
	})
}
