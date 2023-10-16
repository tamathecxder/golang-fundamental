package structs

import "fmt"

type Car struct {
	Brand string
	Speed int
}

func (car Car) toggleNOS(status bool) {
	if status {
		fmt.Println(car.Brand, "NOS is active")
	} else {
		fmt.Println(car.Brand, "NOS is inactive")
	}
}

func StructMethods() {
	mercedes := Car{
		Brand: "Mercedes",
		Speed: 300,
	}

	mercedes.toggleNOS(true)
}
