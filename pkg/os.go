package pkg

import (
	"fmt"
	"os"
)

func PkgOS() {
	// args := os.Args

	// fmt.Println(args)
	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hostname:", hostname)
	}
}
