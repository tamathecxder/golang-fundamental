package datatypes

import "fmt"

func Variables() {
	var username string
	var pi = 3.14

	username = "tamathecxder"

	fmt.Println(username)
	fmt.Println(pi)

	// without var
	language := "Sunda"
	language = "Indonesia"

	fmt.Println(language)

	// multiple declaration
	var (
		oshi = "Indira Seruni"
		age  = 19
	)

	fmt.Println(oshi, age)
}
