package pointer

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

	fmt.Println("Inside Method:", man.Name)
}

func PointerMethod() {
	tama := Man{"Tama"}
	tama.Married()

	fmt.Println("Outside Method:", tama.Name)
}
