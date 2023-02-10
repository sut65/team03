package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

// ============================================================= Not Blank
func TestBookingValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check start not blank", func(t *testing.T) {
		booking := Booking{
			BranchID: &test,
			// Start:       time.Now(), //// ผิด ไม่ใส่ Start
			Stop:       time.Now().AddDate(0, 0, 1),
			RoomID:     &test,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกเวลาเริ่มเข้าพัก"))
	})

	t.Run("check stop not blank", func(t *testing.T) {
		booking := Booking{
			BranchID: &test,
			Start:    time.Now(),
			// Stop:       time.Now().AddDate(0, 0, 1), //// ผิด ไม่ใส่ Stop
			RoomID:     &test,
			CustomerID: &test,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกวันที่สิ้นสุดการพัก"))
	})

	t.Run("check branch is not nil", func(t *testing.T) {
		booking := Booking{
			BranchID:   nil, // ผิด ไม่ใส่ BranchID
			Start:      time.Now(),
			Stop:       time.Now().AddDate(0, 0, 2),
			RoomID:     &test,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกสาขาของโรงแรม"))
	})

	t.Run("check room is not nil", func(t *testing.T) {
		booking := Booking{
			BranchID:   &test,
			Start:      time.Now(),
			Stop:       time.Now().AddDate(0, 0, 2),
			RoomID:     nil, // ผิด ไม่ใส่ RoomID
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกห้องพัก"))
	})

	t.Run("check customer is not nil", func(t *testing.T) {
		booking := Booking{
			BranchID:   &test,
			Start:      time.Now(),
			Stop:       time.Now().AddDate(0, 0, 2),
			RoomID:     &test,
			CustomerID: nil, // ผิด ไม่ใส่ CustomerID
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาเข้าสู่ระบบ"))
	})
}

// ============================================================= Date Now ->
func TestValidateDateMustBeFutureOrNow(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check start must be after now and be now", func(t *testing.T) {
		booking := Booking{
			BranchID:   &test,
			Start:      time.Now().AddDate(0, 0, -1), //ผิด เป็นอดีต
			Stop:       time.Now().AddDate(0, 0, 1),
			RoomID:     &test,
			CustomerID: &test,
		}

		ok, err := govalidator.ValidateStruct(booking)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("เวลาในการเข้าพักไม่ถูกต้อง(ห้ามเป็นอดีต)"))
	})
}

// ============================================================= Stop = Start+1 ->
func TestValidateStopMustbeAfterStartOneDay(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check stop must be after start around one day", func(t *testing.T) {
		fixtures := []time.Time{
			time.Now(),                  //ผิด เพราะมาก่อนวันเข้าพัก
			time.Now().AddDate(0, 0, 2), //ผิด เพราะตรงกับวันเข้าพัก
		}

		for _, fixture := range fixtures {
			booking := Booking{
				BranchID:   &test,
				Start:      time.Now().AddDate(0, 0, 2),
				Stop:       fixture, //ผิด
				RoomID:     &test,
				CustomerID: &test,
			}

			ok, err := govalidator.ValidateStruct(booking)

			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("เวลาสิ้นสุดการพักต้องอยู่หลังวันเข้าพักอย่างน้อย 1 วัน"))
		}
	})
}
