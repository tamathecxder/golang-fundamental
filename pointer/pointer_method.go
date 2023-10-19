package pointer

import "fmt"

type NewAddress struct {
	City, Province, Country string
}

func fillCountry(address *NewAddress, countryName string) {
	address.Country = countryName
}

func PointerMethod() {
	myAddress := NewAddress{
		City:     "Cianjur",
		Province: "Jawa Barat",
	}

	var myAddressPointer *NewAddress = &myAddress

	fillCountry(myAddressPointer, "Netherland")

	fmt.Println(myAddress)
	fmt.Println(myAddressPointer)
}
