package pkg

import (
	"flag"
	"fmt"
)

func PkgFlag() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "admin", "Put your database username")
	password := flag.String("password", "admin", "Put your database password")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
}
