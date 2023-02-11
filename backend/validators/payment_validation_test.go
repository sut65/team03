package validators

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
)

func TestPaymentValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check Customer", func(t *testing.T) {
		p := Payment{
			CustomerID:      0,
			PaymentMethodID: 1,
			MethodID:        1,
			PlaceID:         1,
			Time:            time.Now(),
			Picture:         "data:image/png;base64,iVBORIZIgAAAAJcEh",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please Login"))
	})
	t.Run("Check PaymentMethod", func(t *testing.T) {
		p := Payment{
			CustomerID:      1,
			PaymentMethodID: 0,
			MethodID:        1,
			PlaceID:         1,
			Time:            time.Now(),
			Picture:         "data:image/png;base64,iVBORIZIgAAAAJcEh",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose payment method"))
	})
	t.Run("Check Method", func(t *testing.T) {
		p := Payment{
			CustomerID:      1,
			PaymentMethodID: 1,
			MethodID:        0,
			PlaceID:         1,
			Time:            time.Now(),
			Picture:         "data:image/png;base64,iVBORIZIgAAAAJcEh",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose method"))
	})
	t.Run("Check Place", func(t *testing.T) {
		p := Payment{
			CustomerID:      1,
			PaymentMethodID: 1,
			MethodID:        1,
			PlaceID:         0,
			Time:            time.Now(),
			Picture:         "data:image/png;base64,iVBORIZIgAAAAJcEh",
		}
		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Choose place"))
	})

}
