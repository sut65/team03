package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

func TestCHK_PaymentValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check payment not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           nil, // ผิด ไม่ใส่ PaymentID
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกรายการชำระเงิน"))
	})

	t.Run("check Check payment status not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: nil, // ผิด ไม่ใส่ CHK_PaymentID
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}
		ok, err := govalidator.ValidateStruct(chkp)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกสถานะการชำระเงิน"))
	})

	t.Run("check date time is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Time{}, // ผิด ไม่ใส่ Date_time
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาเลือกวันที่ที่ทำการชำระเงิน"))
	})

	t.Run("check amount is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			// Amount:              17000, // ผิด ไม่ใส่ Amount
			Description: "-",
			EmployeeID:  &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาระบุจำนวนเงิน"))
	})

	t.Run("check employee is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          nil, // ผิด ไม่ใส่ EmployeeID
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณาเข้าสู่ระบบ"))
	})
}

func TestValidateAmountPositive(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check amount must be positive", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              -17000, //ผิด เงินติดลบ
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("จำนวนเงินไม่สามารถติดลบได้"))
	})
}

func TestValidateDateTimeIsPast(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)
	t.Run("check amount must be positive", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2024, 2, 7, 0, 0, 0, 0, time.UTC), //ผิด เป็นอนาคต
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("เวลาในการชำระเงินไม่ถูกต้อง(ห้ามเป็นอนาคต)"))
	})
}
