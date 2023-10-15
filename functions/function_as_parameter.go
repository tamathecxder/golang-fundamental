package functions

import "fmt"

type Filter func(string) string

func sendHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)

	fmt.Println("Hello,", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func FunctionAsParameters() {
	user1Name := "Asep"
	sendHelloWithFilter(user1Name, spamFilter)
}
