package loops

import "fmt"

func BreakKeyword() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("i-", i)
	}
}

func ContinueKeyword() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Index:", i)
	}
}
