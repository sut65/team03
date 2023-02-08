package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestReviewPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	re := Review{
		Comment: "ประทับใจทุกอย่างของโรงแรม",
		Star: 5,
		Reviewdate: time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(re)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(gomega.BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(gomega.BeNil())
}

func TestReviewValidateNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Comment not blank", func(t *testing.T) {
		//Comment ห้ามว่าง
		re := Review{
			Comment: "",
			Star: 5,
			Reviewdate: time.Now(),
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(re)
		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(gomega.BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(gomega.BeNil())
		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(gomega.Equal("Comment not blank"))
	})

}

func TestReviewMustBeValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	test_data := ""
	for i := 0; i < 201; i++ {
		test_data += "A"
	}

	// ข้อมูลถูกต้องบาง field
	re := Review{
		Comment: test_data,
		Star: 5,
		Reviewdate: time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(re)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Comment length must be between 0 - 200"))
}