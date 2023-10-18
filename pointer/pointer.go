package pointer

import "fmt"

type Address struct {
	City, Province, Country string
}

func Pointer() {
	address1 := Address{
		City:     "Cianjur",
		Province: "West Java",
		Country:  "Indonesia",
	}

	address2 := &address1
	address3 := &address1

	address2.City = "Sumedang"

	*address2 = Address{
		City:     "Medan",
		Province: "North Sumatera",
		Country:  "Indonesia",
	}

	var address4 *Address = new(Address)

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}

// NOTE:
// & pada saat assign variable itu digunakan untuk menjadikan variable tersebut pointer atau
// mereference variable tersebut agar menjadi pass by reference

// pada contoh diatas, jika misal address2 ingin memisahkan reference nya agar dia reinitialize
// nilai baru, maka gunakan & juga saat reinit nilai nya, ini menyebabkan didalam memori
// yang sama tiap variable nya akan memiliki reference masing masing namun dalam lingkup 1 memori yang sama

// pada pointer, jika variable reference ingin diubah juga referensinya agar mengikuti pointer nya
// saat reinitialize nilai pointer, maka bisa menggunakan *{nama_pointer} agar
// reference bisa mengikuti nilai baru saat varialbe pointer melakukan re-init, konteksnya masih dalam memori yang sama
