package functions

import "fmt"

func Closure() {
	counter := 0

	increment := func() {
		fmt.Println("Increment in running")
		counter++
	}

	increment()
	increment()
	fmt.Println("Counter", counter)
}
