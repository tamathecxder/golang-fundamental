package functions

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Agus"
	lastName = "Bakrie"

	return
}

func NamedReturns() {
	firstName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}
