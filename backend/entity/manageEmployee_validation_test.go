package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeeValidateNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Name not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "", //ผิด
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name not blank"))
	})

	t.Run("Check Email not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "Sobsa tugwan",
			Email:        "", //ผิด
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Email not blank"))
	})

	t.Run("Check Username not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "",//ผิด
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Username not blank"))
	})

	t.Run("Check Phonenumber not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  "",//ผิด
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Tel not blank"))
	})

	t.Run("Check Address not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "",//ผิด
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Address not blank"))
	})

	t.Run("Check Password not blank", func(t *testing.T) {

		e := Employee{
			PersonalID:   "1104200258432",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "",//ผิด
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Password not blank"))
	})
}

func TestEmployeeCheckPersonalid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	fixtures := []string{
		"A0000000000000", //ผิดเพราะมีตัวอักษร
		"125224540000",   // ผิดเพราะเลขไม่ครบ 13
		"b1522a22552250", //ผิดเพราะมีตัวอักษร
		"sdhsfgjfjijdfj",
	}
	for _, fixture := range fixtures {
		e := Employee{
			PersonalID:   fixture, // ผิด
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("PersonalID is not vaild"))
	}

}

func TestEmployeeEmailMustBeValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	e := Employee{
		PersonalID:   "1104200258432",
		Employeename: "Sobsa tugwan",
		Email:        "fgss", //ผิด
		Eusername:    "ESobsa",
		Password:     "Sobsa01",
		Phonenumber:  "0905452001",
		Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
	}

	ok, err := govalidator.ValidateStruct(e)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("Email is not vaild"))

}

func TestEmployeetCheckPhonenumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	fixtures := []string{
		"190545200", // ผิดเพราะเลขไม่ครบ 10 และตัวนำหน้าไม่ใช่ 0
		"09548125", // ผิดเพราะเลขไม่ครบ 10
		"b288525", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
		"b28g525b", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
	}
	for _, fixture := range fixtures{
		e := Employee{
			PersonalID:   "1104200258430",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     "Sobsa01",
			Phonenumber:  fixture,
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Phonenumber is not vaild"))
	}
}

func TestEmployeetUsername(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	fixtures := []string{
		"E0258151", //คำนำหน้าถูก แต่มีตัวเลข
		"esdfssd", // คำนำหน้าถูก แต่เป็นพิมพ์เล็ก และอักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		"bdfdf", //คำนำหน้าผิด แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		"Esfdfdf", //คำนำหน้าถูก แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
	}
	for _, fixture := range fixtures{
		e := Employee{
			PersonalID:   "1104200258430",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    fixture,
			Password:     "Sobsa01",
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Username must be is Begin with E and The second letter must start with A-Z and must not number"))
	}
}

func TestEmployeetPassword(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	fixtures := []string{
		"bdfdf", //คำนำหน้าผิด แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		"Esf", //คำนำหน้าถูก แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
	}
	for _, fixture := range fixtures{
		e := Employee{
			PersonalID:   "1104200258430",
			Employeename: "Sobsa tugwan",
			Email:        "Sobsa@gmail.com",
			Eusername:    "ESobsa",
			Password:     fixture,
			Phonenumber:  "0905452001",
			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Password must be more than 6 characters"))
	}
}