package validators

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"
) 



func TestCustomerEmailisNotValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	test := uint(1)

	//ทำการตรวจสอบ Email ต้องถูกต้องมา Pattern ของ Email
	cus := Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	30,
			Phone: "0985441236",
			Password: "SD123456",
			Email: "",
			Nametitle_ID: &test,
			Gender_ID: 	&test,
			Province_ID: &test,	
	}

	ok, err := govalidator.ValidateStruct(cus)

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
			Password: "SD123456",
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
			Password: "SD123456",
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
			Password: "SD123456",
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
			Password: "SD123456",
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

func TestCustomerPassword(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	test := uint(1)
	//ทำการตรวจสอบ Password ต้องมากกว่า 6 หรือเท่ากับ 6 ตัว
	fixtures := []string{
		"abcde", //คำนำหน้าผิด แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		"Esf", //คำนำหน้าถูก แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
	}
	for _, fixture := range fixtures{
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	20,
			Password: fixture,
			Phone: "0985441236",
			Email: "S@gmail.com",
			Nametitle_ID: &test,
			Gender_ID: 	&test,
			Province_ID: &test,
		}

		ok, err := govalidator.ValidateStruct(cus)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Password must be more than or equal to 8 characters")) 
	}

	t.Run("Check Password", func(t *testing.T) {

		fixtures := []string{
			"bdfdf453", //คำนำหน้าผิด แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
			"5532323sf", //คำนำหน้าถูก แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		}
		//Address ห้ามว่าง
		for _, fixture := range fixtures{
			cus:= Customer{
				FirstName: "Sandee",
				LastName: "Memak",
				Age: 	20,
				Password: fixture,
				Phone: "0985441236",
				Email: "S@gmail.com",
				Nametitle_ID: &test,
				Gender_ID: 	&test,
				Province_ID: &test,
			}

		ok, err := govalidator.ValidateStruct(cus)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Password must contain at least 1 character A-Z."))
	}
})
}

func TestCustomerCheckPhone(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	test := uint(1)
	//ทำการตรวจสอบ Phonenumber ต้องมีตัวเลข 0-9 เท่ากับ 10 ตัว ไม่มีตัวอักษร และ ขึ้นต้นด้วยเลข 0
	fixtures := []string{
		"190545200", // ผิดเพราะเลขไม่ครบ 10 และตัวนำหน้าไม่ใช่ 0
		"09548125", // ผิดเพราะเลขไม่ครบ 10
		"b288525", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
		"b28g525b", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
	}
	for _, fixture := range fixtures{
		cus:= Customer{
			FirstName: "Sandee",
			LastName: "Memak",
			Age: 	20,
			Password: "SD123456",
			Phone: fixture,
			Email: "S@gmail.com",
			Nametitle_ID: &test,
			Gender_ID: 	&test,
			Province_ID: &test,
		}

		ok, err := govalidator.ValidateStruct(cus)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Phonenumber is not valid"))
	}
}
