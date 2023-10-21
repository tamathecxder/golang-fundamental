package pkg

import (
	"fmt"
	"strings"
)

func PkgStrings() {
	var name string = "Yudistira Eka Pratama"
	var desc string = "I am an anaconda   "

	fmt.Println(strings.Contains(name, "Eka"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Trim(desc, " "))
	fmt.Println(strings.ReplaceAll(name, "Eka", "Tehyung Jimin Tehyung"))
}
