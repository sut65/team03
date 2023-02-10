package entity

import (
	"testing"
	"time"

	//"github.com/asaskevich/govalidator"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCheckInOutValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check check-in time not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Time{},
			CheckOutTime:       time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select check-in time"))
	})

	t.Run("check check-out time not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          &test,
			CheckInTime:        time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			CheckOutTime:       time.Time{},
			CheckInOutStatusID: &test,
			EmployeeID:         &test,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select check-out time"))
	})

	t.Run("check booking not blank", func(t *testing.T) {
		cio := CheckInOut{
			BookingID:          nil,
			CheckInTime:        time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			CheckOutTime:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
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
			CheckInTime:        time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			CheckOutTime:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
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
			CheckInTime:        time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			CheckOutTime:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			CheckInOutStatusID: &test,
			EmployeeID:         nil,
		}

		ok, err := govalidator.ValidateStruct(cio)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Signin"))
	})
}
