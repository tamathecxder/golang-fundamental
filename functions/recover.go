package functions

import "fmt"

func panicTest(success bool) {
	if !success {
		panic("Application is stopped unexpectedly becasue bla..bla.bla")
	}

	message := recover()

	if message != nil {
		fmt.Println("Error Trace:", message)
	}
}

func runApplication() {
	defer panicTest(false)

	fmt.Println("App is running on xxxxx")
}

func Recover() {
	runApplication()
}
