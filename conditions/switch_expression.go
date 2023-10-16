package datatypes

import "fmt"

func SwitchExpressions() {
	name := "Jihyo"
	age := 21

	switch name {
	case "Mina":
		fmt.Println("Good Morning, My Wife, Mina!")
	case "Dahyun":
		fmt.Println("Good Morning, My Wife, Dahyun!")
	case "Jihyo":
		fmt.Println("Good Morning, My Wife, Jihyo!")
	case "Tzuyu":
		fmt.Println("Good Morning, My Wife, Tzuyu!")
	default:
		fmt.Println("...")
		fmt.Println("I need to take pills!")
	}

	// short expression
	switch length := len(name); length != 0 {
	case true:
		fmt.Println("Valid")
	case false:
		fmt.Println("Invalid")
	}

	// without condition
	switch {
	case age >= 21:
		fmt.Println("Adult")
	case age >= 18:
		fmt.Println("Young Adult")
	case age < 18:
		fmt.Println("Young")
	default:
		fmt.Println("Error")
	}
}
