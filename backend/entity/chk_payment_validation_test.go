package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของการจองห้องพักแล้วถูกทั้งหมด
func TestCHK_PaymentValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check payment not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           nil,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select Payment"))
	})

	t.Run("check Check payment status not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: nil,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}
		ok, err := govalidator.ValidateStruct(chkp)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please select Status"))
	})

	t.Run("check date time is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Time{},
			Amount:              17000,
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select date time"))
	})

	t.Run("check amount is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Description:         "-",
			EmployeeID:          &test,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please input Amount"))
	})

	t.Run("check employee is not blank", func(t *testing.T) {
		chkp := CHK_Payment{
			PaymentID:           &test,
			CHK_PaymentStatusID: &test,
			Date_time:           time.Date(2023, 2, 7, 0, 0, 0, 0, time.UTC),
			Amount:              17000,
			Description:         "-",
			EmployeeID:          nil,
		}

		ok, err := govalidator.ValidateStruct(chkp)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Signin"))
	})
}

func (chkp CHK_Payment) ValidateDateTimeNotFuture() error {
	today := time.Now().UTC()
	if chkp.Date_time.After(today) {
		return fmt.Errorf("Date time must not be the future")
	}
	return nil
}

func TestValidateDateTimeNotFuture(t *testing.T) {
	test := uint(1)

	chkp := CHK_Payment{
		PaymentID:           &test,
		CHK_PaymentStatusID: &test,
		Date_time:           time.Date(2024, 2, 7, 0, 0, 0, 0, time.UTC),
		Amount:              17000,
		Description:         "-",
		EmployeeID:          &test,
	}
	err := chkp.ValidateDateTimeNotFuture()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Validation succeeded")
	}
}
