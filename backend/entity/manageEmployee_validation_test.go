package entity

import (

	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

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

}

func TestEmployeeCheckPersonalid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	fixtures := []string{
		"A0000000000000", //ผิดเพราะมีตัวอักษร
		"125224540000", // ผิดเพราะเลขไม่ครบ 13
		"b1522a22552250", //ผิดเพราะมีตัวอักษร
		"sdhsfgjfjijdfj",
	}
	for _, fixture := range fixtures{
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

// func TestEmployeeEmailMustBeValid(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 		e := Employee{
// 			PersonalID:   "1104200258432",
// 			Employeename: "Sobsa tugwan",
// 			Email:        "fgss", //ผิด
// 			Eusername:    "ESobsa",
// 			Password:     "Sobsa01",
// 			Phonenumber:  "0905452001",
// 			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
// 		}

// 		ok, err := govalidator.ValidateStruct(e)

// 		g.Expect(ok).NotTo(gomega.BeTrue())

// 		g.Expect(err).ToNot(gomega.BeNil())

// 		g.Expect(err.Error()).To(gomega.Equal("Email is not vaild"))

// }

// func TestEmployeetCheckPhonenumber(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)
// 	fixtures := []string{
// 		"190545200", // ผิดเพราะเลขไม่ครบ 10 และตัวนำหน้าไม่ใช่ 0
// 		"09548125", // ผิดเพราะเลขไม่ครบ 10
// 		"b288525", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
// 		"b28g525b", //ผิดเพราะมีตัวอักษร เลขไม่ครบ 10
// 	}
// 	for _, fixture := range fixtures{
// 		e := Employee{
// 			PersonalID:   "1104200258430", 
// 			Employeename: "Sobsa tugwan",
// 			Email:        "Sobsa@gmail.com",
// 			Eusername:    "ESobsa",
// 			Password:     "Sobsa01",
// 			Phonenumber:  fixture,
// 			Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
// 		}

// 		ok, err := govalidator.ValidateStruct(e)

// 		g.Expect(ok).NotTo(gomega.BeTrue())

// 		g.Expect(err).ToNot(gomega.BeNil())

// 		g.Expect(err.Error()).To(gomega.Equal("Phonenumber is not vaild"))
// 	}
// }