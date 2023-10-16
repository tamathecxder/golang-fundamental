package structs

import "fmt"

type Motorcycle struct {
	Brand string
	Speed int
}

type HasMotorcycleInfo interface {
	getMotorcycleInfo() string
}

func spillMotorcycleBrand(hasMotorcycleInfo HasMotorcycleInfo) {
	fmt.Println("This is", hasMotorcycleInfo.getMotorcycleInfo())
}

func (Motorcycle Motorcycle) getMotorcycleInfo() string {
	return Motorcycle.Brand
}

func (Motorcycle Motorcycle) toggleNOS(status bool) {
	if status {
		fmt.Println(Motorcycle.Brand, "NOS is active")
	} else {
		fmt.Println(Motorcycle.Brand, "NOS is inactive")
	}
}

func Interfaces() {
	mercedes := Motorcycle{
		Brand: "Kawasaki",
		Speed: 300,
	}

	mercedes.toggleNOS(true)
	spillMotorcycleBrand(mercedes)
}
