package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestServiceValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check Customer", func(t *testing.T) {
		s := Service{
			CustomerID:    0,
			Time:          time.Now(),
			FoodID:        1,
			DrinkID:       1,
			AccessoriesID: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please login"))
	})
	t.Run("Check Food", func(t *testing.T) {
		s := Service{
			CustomerID:    1,
			Time:          time.Now(),
			FoodID:        0,
			DrinkID:       1,
			AccessoriesID: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Choose some order of Food"))
	})
	t.Run("Check Drink", func(t *testing.T) {
		s := Service{
			CustomerID:    1,
			Time:          time.Now(),
			FoodID:        1,
			DrinkID:       0,
			AccessoriesID: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Choose some order of Drink"))
	})
	t.Run("Check Accessories", func(t *testing.T) {
		s := Service{
			CustomerID:    1,
			Time:          time.Now(),
			FoodID:        1,
			DrinkID:       1,
			AccessoriesID: 0,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Choose some order of Accessories"))
	})
}
