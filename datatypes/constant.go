package datatypes

import "fmt"

func Constants() {
	const cityName string = "Cianjur"
	const year uint16 = 2023

	const (
		isMuslim = true
		isSingle = false
	)

	fmt.Println(cityName, year)
	fmt.Println(isMuslim, isSingle)
}
