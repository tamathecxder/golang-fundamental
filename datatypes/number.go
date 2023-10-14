package datatypes

import "fmt"

func Numbers() {
	/*
		Integer
		- int8 	= min: -128, 					max: 127
		- int16 = min: -32788, 					max: 32767
		- int32 = min: -2147483648, 			max: 2147483647
		- int64 = min: -9223372036854775808, 	max: 9223372036854775807
	*/

	fmt.Println("One = ", 1)
	fmt.Println("Two = ", 2)
	fmt.Println("Three Point 5 = ", 3.5)

	/*
			Unsigned Integer
			- uint8  = min: 0,   max: 255
			- uint16 = min: 0,	max: 65535
			- uint32 = min: 0,	max: 4294967295
			- uint64 = min: 0, 	max: 1844_ _ _ _  _ _ _ _ _ _ _  _ _ _ _
		/*
			Unsigned Integer
			- float32 	= min: 1.18x10 ^-38		max: 1.18x10 ^38
			- uint64 	= min: 2.23x10 ^-308	max: 2.23x10 ^308
	*/

	/*
		Alias
		- byte 		= uint8
		- rune		= int32
		- int		= minimal int32
		- uint		= minimal uint32
	*/
}
