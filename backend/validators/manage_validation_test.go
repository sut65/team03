package validators

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	. "github.com/sut65/team03/entity"

)

func TestEmployeePass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	e := Employee{
		PersonalID:   "1104200258432",
		Employeename: "Sobsa tugwan",
		Email:        "Sobsa@gmail.com",
		Eusername:    "ESobsa",
		Password:     "Sobsa01",
		Phonenumber:  "0905452001",
		Address:      "219 m.10, nongprajak s, nongsham d, Ayutthaya 13000",
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(e)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(gomega.BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(gomega.BeNil())
}

func TestEmployeeValidateNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Name not blank", func(t *testing.T) {
		//Name ห้ามว่าง
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
		//Email ห้ามว่าง
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
		//Username ห้ามว่าง
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
		//Tel ห้ามว่าง
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
		//Address ห้ามว่าง
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
		//Password ห้ามว่าง
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

		//ทำการตรวจสอบ Personalid ต้องมีตัวเลข 0-9 เท่ากับ 13 ตัว ไม่มีตัวอักษร
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

		g.Expect(err.Error()).To(gomega.Equal("PersonalID is not valid"))
	}

}

func TestEmployeeEmailMustBeValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	//ทำการตรวจสอบ Email ต้องถูกต้องมา Pattern ของ Email
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

	g.Expect(err.Error()).To(gomega.Equal("Email is not valid"))

}

func TestEmployeetCheckPhonenumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	//ทำการตรวจสอบ Phonenumber ต้องมีตัวเลข 0-9 เท่ากับ 10 ตัว ไม่มีตัวอักษร และ ขึ้นต้นด้วยเลข 0
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

		g.Expect(err.Error()).To(gomega.Equal("Phonenumber is not valid"))
	}
}

func TestEmployeetUsername(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

		//ทำการตรวจสอบ Username ต้องขึ้นต้นด้วยอักษร E พิมพ์ใหญ่เท่านั้น และตัวอักษรตัวที่สองต้องเป็นพิมพ์ใหญ่ ตัวต่อไปเป็น a-z และ A-Z ไม่จำกัด ยกเว้นตัวเลขเท่านั้น
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

	//ทำการตรวจสอบ Password ต้องมากกว่า 6 หรือเท่ากับ 6 ตัว
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

		g.Expect(err.Error()).To(gomega.Equal("Password must be more than or equal to 6 characters"))
	}

	t.Run("Check Password", func(t *testing.T) {

		fixtures := []string{
			"bdfdf453", //คำนำหน้าผิด แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
			"5532323sf", //คำนำหน้าถูก แต่อักษรตัวที่สองเป็นพิมพ์เล็ก ซึ่งผิด
		}
		//Address ห้ามว่าง
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

		g.Expect(err.Error()).To(gomega.Equal("Password must contain at least 1 character A-Z."))
	}
})


}