package structs

import (
	"errors"
	"fmt"
)

func divide(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("can be divided by zero")
	} else {
		return value / divider, nil
	}
}

func ErrorInterfaces() {
	result, err := divide(12, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
