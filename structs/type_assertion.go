package structs

import "fmt"

func random() interface{} {
	return "bipbopbipbop"
}

func TypeAssertions() {
	var data interface{} = random()
	var converted string = data.(string)

	fmt.Println(converted)

	// more safety conversion
	switch value := data.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is integer")
	default:
		fmt.Println("Unknown Type")
	}
}
