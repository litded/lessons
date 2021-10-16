package main

import "fmt"

type SalaryCalculator interface {
	CalcSalary() int
}

type Intern struct {
	coeff int
	base  int
}

func (i Intern) CalcSalary() int {
	return i.base + i.coeff - 1000
}

type Permanent struct {
	salary    int
	premiales int
}

func (p Permanent) CalcSalary() int {
	return p.salary + p.premiales
}

type Contract struct {
	fixSalary int
}

func (c Contract) CalcSalary() int {
	return c.fixSalary
}

type Freeelancer struct {
	perrHour int
	hours    int
}

func (f Freeelancer) CalcSalary() int {
	return f.hours * f.perrHour
}

func TotalSalary(salary []SalaryCalculator) {
	total := 0
	for _, s := range salary {
		total += s.CalcSalary()
	}
	fmt.Println("You need : ", total, "DENEG")
}

func main() {
	perm1 := Permanent{10000, 300}
	perm2 := Permanent{12000, 350}
	fl1 := Freeelancer{10, 75}
	fl2 := Freeelancer{12, 80}
	cntr1 := Contract{7500}

	intrn1 := Intern{25, 500}
	employess := []SalaryCalculator{perm1, perm2, fl1, fl2, cntr1, intrn1}
	TotalSalary(employess)
}
