package pkg

import (
	"fmt"
	"strconv"
)

func PkgStrconv() {
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(boolean)
	}

	numberString, err := strconv.ParseInt("1000000", 10, 64)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(numberString)
	}

	numberFloat := strconv.FormatInt(1000000, 2)
	fmt.Println(numberFloat)

	convToInt := 500
	convToStr := "10"

	fmt.Println(strconv.Itoa(convToInt))
	fmt.Println(strconv.Atoi(convToStr))
}
