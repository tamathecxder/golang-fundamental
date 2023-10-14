package datatypes

import "fmt"

func TypeDeclarations() {
	type NoKTP string

	var ktpMamat NoKTP = "123412341234"

	fmt.Println(ktpMamat)
}
