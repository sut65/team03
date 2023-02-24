package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

func TestCheckroomValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check status is not nil", func(t *testing.T) {
		cr := Checkroom{
			StatusID:   nil,
			RoomID:     &test,
			ProductID:  &test,
			DamageID:   &test,
			Date:       time.Now(),
			EmployeeID: &test,
		}

		ok, err := govalidator.ValidateStruct(cr)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Status"))
	})

	t.Run("check damage is not nil", func(t *testing.T) {
		cr := Checkroom{
			StatusID:   &test,
			RoomID:     nil,
			ProductID:  &test,
			DamageID:   &test,
			Date:       time.Now(),
			EmployeeID: &test,
		}

		ok, err := govalidator.ValidateStruct(cr)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Room"))
	})

	t.Run("check room is not nil", func(t *testing.T) {
		cr := Checkroom{
			StatusID:   &test,
			RoomID:     nil,
			ProductID:  &test,
			DamageID:   &test,
			Date:       time.Now(),
			EmployeeID: &test,
		}

		ok, err := govalidator.ValidateStruct(cr)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Room"))
	})

}

