package functions

import "fmt"

func logging() {
	fmt.Println("logging: xxxxx is executed")
}

func divide(value int) int {
	return 10 / value
}

func runApp() {
	// - error -
	// fmt.Println("App is running on bla..bla..bla")
	// result := divide(0)
	// fmt.Println("result:", result)
	// logging()

	defer logging()
	fmt.Println("App is running on bla..bla..bla")
	result := divide(10)
	fmt.Println("result:", result)
}

func Defer() {
	runApp()
}
