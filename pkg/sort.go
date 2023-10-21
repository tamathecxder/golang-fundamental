package pkg

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	Salary int
}

type EmployeeSlice []Employee

func (employee EmployeeSlice) Len() int {
	return len(employee)
}

func (employee EmployeeSlice) Less(i, j int) bool {
	return employee[i].Salary < employee[j].Salary
}

func (employee EmployeeSlice) Swap(i, j int) {
	employee[i], employee[j] = employee[j], employee[i]
}

func PkgSort() {
	employees := []Employee{
		{"Gilang", 1500000},
		{"Asep", 10000000},
		{"Gilang", 50000},
		{"Siti", 7000000},
		{"Dinda", 1200000},
	}

	fmt.Println(employees)

	sort.Sort(EmployeeSlice(employees))

	fmt.Println(employees)

	// for _, emp := range employees {
	// 	fmt.Println(emp.Name)
	// }

}
