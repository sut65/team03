package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
) 


func TestCustomerSuccess(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	cu := Customer{
		FirstName: "Sandee",
			LastName: "Memak",
			Age: 	30,
			Phone: "0985441236",
			Email: "S@gmail.com",	
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(cu)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(gomega.BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(gomega.BeNil())
}


func TestCustomerEmailisNotValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	//ทำการตรวจสอบ Email ต้องถูกต้องมา Pattern ของ Email
	e := Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	30,
			Phone: "0985441236",
			Email: "",	
	}

	ok, err := govalidator.ValidateStruct(e)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("Email is not valid"))

}



func TestCustomerValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	test := uint(1)

	t.Run("check age is not negative number", func(t *testing.T) {
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	-20,
			Phone: "0985441236",
			Email: "S@gmail.com",
			Nametitle_ID: &test,
			Gender_ID: 	&test,
			Province_ID: &test,
		
		}

	ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("อายุไม่สามารถติดลบได้"))
	})


	t.Run("check nametitle is not nil", func(t *testing.T) {
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	20,
			Phone: "0985441236",
			Email: "S@gmail.com",
			Nametitle_ID: nil,
			Gender_ID: 	&test,
			Province_ID: &test,
		}

	ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Nametitle"))
	})

	t.Run("check gender is not nil", func(t *testing.T) {
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	20,
			Phone: "0985441236",
			Email: "S@gmail.com",
			Nametitle_ID: &test,
			Gender_ID: 	nil,
			Province_ID: &test,
		}

	ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Gender"))
	})

	t.Run("check province is not nil", func(t *testing.T) {
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	20,
			Phone: "0985441236",
			Email: "S@gmail.com",
			Nametitle_ID: &test,
			Gender_ID: 	&test,
			Province_ID: nil,
		}

	ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Please select Province"))
	})


}