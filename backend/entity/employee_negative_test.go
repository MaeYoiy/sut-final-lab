package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func EmployeeTestBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	e := Employee{
		Name:       "",
		Email:      "Sobsa@gmail.com",
		EmployeeID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Name is not blank"))
}
