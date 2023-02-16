package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func EmployeeTest(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	e := Employee{
		Name:       "Thansiri",
		Email:      "Sobsa@gmail.com",
		EmployeeID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}
