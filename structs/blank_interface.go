package structs

import "fmt"

func sameAsAny(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return "default"
	} else {
		return true
	}
}

func BlankInterfaces() {
	test := sameAsAny(1)

	fmt.Println(test)
}
