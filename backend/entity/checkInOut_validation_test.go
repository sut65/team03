package entity

import (
	"testing"
	"time"

	//"github.com/asaskevich/govalidator"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookingNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	var Sobsa Employee
	db.Raw("SELECT * FROM employees WHERE employeename  = ?", "Sobsa tugwan").Scan(&Sobsa)

	var checkout CheckInOutStatus
	db.Raw("SELECT * FROM Check_In_Out_Statuses WHERE id = ?", "1").Scan(&checkout)

	var n uint = 1
	var val *uint = &n
	var val2 *uint = &n

	cio := CheckInOut{
		BookingID:          nil,
		CheckInTime:        time.Now(),
		CheckOutTime:       time.Now(),
		CheckInOutStatusID: val,
		EmployeeID:         val2,
	}

	ok, err := govalidator.ValidateStruct(cio)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Booking Not Blank"))
}
