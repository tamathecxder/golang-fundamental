package functions

import "fmt"

func stopApp(error bool) {
	if error {
		panic("Application is error")
	}
}

func Panic() {
	fmt.Println("App is running on xxxx")

	stopApp(true)
}
