package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของการจองห้องพักแล้วถูกทั้งหมด
func TestBookingValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check start not blank", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   &test,
			Start:      time.Time{},
			Stop:       time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			RoomID:     &test,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select Start date"))
	})

	t.Run("check stop not blank", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   &test,
			Start:      time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Stop:       time.Time{},
			RoomID:     &test,
			CustomerID: &test,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select Stop date"))
	})

	t.Run("check branch is not nil", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   nil,
			Start:      time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Stop:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			RoomID:     &test,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select branch"))
	})

	t.Run("check room is not nil", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   &test,
			Start:      time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Stop:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			RoomID:     nil,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select room"))
	})

	t.Run("check customer is not nil", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   &test,
			Start:      time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Stop:       time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			RoomID:     &test,
			CustomerID: nil,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Signin"))
	})
}

func (b Booking) ValidateStartBeforeStop() error {
	if b.Start.After(b.Stop) {
		return fmt.Errorf("Start time must be before stop time")
	}
	return nil
}

func (b Booking) ValidateStopAtMostOneYearInFuture() error {
	today := time.Now().UTC()
	maxStop := today.AddDate(0, 0, 2)
	if b.Stop.After(maxStop) {
		return fmt.Errorf("Stop time must be at most one year in the future")
	}
	return nil
}

func TestValidateStartBeforeStop(t *testing.T) {
	t.Run("check start before stop", func(t *testing.T) {
		test := uint(1)
		booking := Booking{
			BranchID:   &test,
			Start:      time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			Stop:       time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			RoomID:     &test,
			CustomerID: &test,
		}
		err := booking.ValidateStartBeforeStop()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Validation succeeded")
		}
	})
}
