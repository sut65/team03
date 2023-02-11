package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
) 


func TestCheckroomValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check status is not nil", func(t *testing.T) {
		cr:= Checkroom{
			StatusID:   nil,
			RoomID:     &test,
			ProductID:  &test,
			DamageID:  &test,
			Date:   time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC),
			EmployeeID: &test,
		}

	ok, err := govalidator.ValidateStruct(cr)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Status"))
	})

	

}
