package functions

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func Variadics() {
	result := sumAll(10, 10, 10, 10, 10)

	fmt.Println(result)

	// with slice argument
	numbers := []int{5, 5, 5, 5, 5}
	total := sumAll(numbers...)

	fmt.Println(total)
}
