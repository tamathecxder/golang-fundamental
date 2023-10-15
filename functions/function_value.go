package functions

import "fmt"

func greetings(name string) string {
	return "Hello, " + name
}

func FunctionValues() {
	greet1 := greetings("Ajeng")

	fmt.Println(greet1)
}
