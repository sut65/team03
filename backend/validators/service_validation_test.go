package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

func TestServiceValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check Customer", func(t *testing.T) {
		s := Service{
			CustomerID:      0,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Login"))
	})
	t.Run("Check Food", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          0,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose Food"))
	})
	t.Run("Check Drink", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         0,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose Drink"))
	})
	t.Run("Check Accessories", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   0,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose Accessories"))
	})
	t.Run("Check Food Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        -1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much food do you want?"))
	})
	t.Run("Check Drink Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       -1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much drink do you want?"))
	})
	t.Run("Check Accessories Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: -1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much accessories do you want?"))
	})
	t.Run("Check Food Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        51,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much food do you want?"))
	})
	t.Run("Check Drink Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       51,
			AccessoriesID:   1,
			AccessoriesItem: 1,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much drink do you want?"))
	})
	t.Run("Check Accessories Item", func(t *testing.T) {
		s := Service{
			CustomerID:      1,
			Time:            time.Now(),
			FoodID:          1,
			FoodItem:        1,
			DrinkID:         1,
			DrinkItem:       1,
			AccessoriesID:   1,
			AccessoriesItem: 51,
		}

		ok, err := govalidator.ValidateStruct(s)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("How much accessories do you want?"))
	})
}
