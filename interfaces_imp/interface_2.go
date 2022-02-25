package main

import "fmt"

type Permanent struct {
	EmpID    int
	BasicPay int
	Pf       int
}

type Contract struct {
	EmpID    int
	BasicPay int
}

type SalaryCalulator interface {
	CalculateSalary() int
}

func (p Permanent) CalculateSalary() int {
	return p.BasicPay + p.Pf
}

func (c Contract) CalculateSalary() int {
	return c.BasicPay
}

func TotalExpenseCalculator(s []SalaryCalulator) {
	TotalExpense := 0
	for _, v := range s {
		TotalExpense += v.CalculateSalary()
	}
	// return TotalExpense
	fmt.Println(TotalExpense)
}
func main() {
	pemp1 := Permanent{
		EmpID:    1,
		BasicPay: 12000,
		Pf:       7000,
	}
	pemp2 := Permanent{
		EmpID:    2,
		BasicPay: 14000,
		Pf:       12000,
	}

	cemp1 := Contract{
		EmpID:    3,
		BasicPay: 6000,
	}
	employees := []SalaryCalulator{pemp1, pemp2, cemp1}
	TotalExpenseCalculator(employees)
}
