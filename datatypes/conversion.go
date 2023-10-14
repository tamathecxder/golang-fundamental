package datatypes

import "fmt"

func Conversions() {
	var _32value int32 = 129
	var _64value int64 = int64(_32value)
	var _8value int8 = int8(_32value)

	fmt.Println(_32value)
	fmt.Println(_64value)
	fmt.Println(_8value)

	var name = "golang"
	var eByte = name[0]
	var eString string = string(eByte)

	fmt.Println(name)
	fmt.Println(eByte)
	fmt.Println(eString)
}
