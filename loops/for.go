package loops

import "fmt"

func ForLoops() {
	// counter := 1

	// with statement
	// for number := 1; number <= 20; number++ {
	// 	if number%2 == 0 {
	// 		fmt.Println("Number: ", number)
	// 	}
	// }

	// without statement
	// for counter <= 10 {
	// 	fmt.Println("Loops: ", counter)
	// 	counter++
	// }

	// iterating collection using regular "for" loops (array, slice, map)
	slice := []string{"Itzy", "Newjeans", "TWICE"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// iterating collection using regular "for range" loops (array, slice, map)
	person := map[string]string{
		"id":   "21",
		"name": "Stefani",
		"age":  "19",
	}

	for _, value := range person {
		// fmt.Println(key, value)
		fmt.Println(value)
	}
}
