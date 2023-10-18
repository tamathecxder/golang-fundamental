package structs

import "fmt"

func storeName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func Nil() {
	var human map[string]string = storeName("Agus")

	if human == nil {
		fmt.Println("Empty data")
	} else {
		fmt.Println(human)
	}
}
