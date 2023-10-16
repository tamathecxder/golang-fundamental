package datatypes

import "fmt"

func ConditionalExpressions() {
	name := "Yudistira"
	age := 19

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 50 {
		fmt.Println("Old")
	} else {
		fmt.Println("Young")
	}

	// short statement
	if length := len(name); length > 10 {
		fmt.Println("The name is too long")
	} else {
		fmt.Println("*Passed")
	}
}
