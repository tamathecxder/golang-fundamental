package datatypes

import "fmt"

func Slices() {
	// var months = [...]string{
	// 	"January",
	// 	"February",
	// 	"March",
	// 	"April",
	// 	"May",
	// 	"June",
	// 	"July",
	// 	"August",
	// 	"September",
	// 	"October",
	// 	"November",
	// 	"December",
	// }

	// var slice1 = months[4:7]
	// fmt.Println(months)

	// fmt.Println(len(slice1)) // get slice length
	// fmt.Println(cap(slice1)) // get slice capacity

	// months[5] = "Arr Change"
	// fmt.Println(slice1)

	// slice1[0] = "April______"
	// fmt.Println(months)

	// var slice2 = months[0:1] // it will change
	// var slice2 = months[11:] // it will not change

	// fmt.Println(slice2)

	// var slice3 = append(slice2, "NANANA")
	// fmt.Println(slice3)

	// slice3[1] = "AOKWOKWOWK"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(months)

	// newSlice := make([]string, 3, 5)

	// newSlice[0] = "Joko"
	// newSlice[1] = "Widodo"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// copySlice := make([]string, len(newSlice), cap(newSlice))
	// copy(copySlice, newSlice)

	// fmt.Println(copySlice)

	// AVOID!
	exampleArray := [...]string{"a", "b", "c"}
	exampleSlice := []string{"x", "y", "z"}

	fmt.Println(exampleArray)
	fmt.Println(exampleSlice)
}
