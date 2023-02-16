package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name is not blank"` //ต้องกรอก
	Email      string `valid:"email"`
	EmployeeID string `valid:"matches(^[LMH]\\d{7}$)~CustomerID is not valid"` // L หรือ  M H มีเลขต่อ 7 ตัว เช่น L1234567
}
