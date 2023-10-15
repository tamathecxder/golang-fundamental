package functions

import "fmt"

func sayThankToWorld(name string) {
	fmt.Println("Thanks, world.", name, "feels so grateful")
}

func Parameters() {
	sayThankToWorld("Agus")
}
