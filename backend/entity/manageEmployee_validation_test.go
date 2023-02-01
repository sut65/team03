package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		e := Employee{
			PersonalID:   1104200258432,
			Employeename: "",
			Email:        "g",
			Eusername:    "tfyt",
			Password:     "tyt",
			Phonenumber:  "fgf",
			Address:      "fdg",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name not blank"))
	})

}
