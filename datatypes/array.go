package datatypes

import "fmt"

func Arrays() {
	var names [3]string
	names[0] = "Yudistira"
	names[1] = "Eka"
	names[2] = "Pratama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1, 2, 3,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}
