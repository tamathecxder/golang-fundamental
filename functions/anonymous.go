package functions

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func blocked(name string) bool {
	switch name {
	case "Coward":
		return true
	}

	return false
}

func Anonymous() {
	user1 := "Ujang"
	user2 := "Coward"

	registerUser(user2, blocked)

	registerUser(user1, func(name string) bool {
		return name == "bitch"
	})
}
