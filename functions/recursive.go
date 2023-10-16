package functions

import "fmt"

func factorialUsingFor(number int) int {
	result := 1

	for i := number; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialUsingRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialUsingRecursive(number-1)
	}
}

func Recursive() {
	loop := factorialUsingFor(3)
	recursive := factorialUsingRecursive(5)

	fmt.Println(loop)
	fmt.Println(recursive)
}
