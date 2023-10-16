package structs

import "fmt"

type User struct {
	Name, Email string
	Age         int
}

func Structs() {
	var azizi User
	azizi.Name = "Azizi Shafaa Asadel"
	azizi.Age = 19
	azizi.Email = "zee@jkt48.com"

	fmt.Println(azizi)
	fmt.Println(azizi.Name)
	fmt.Println(azizi.Age)
	fmt.Println(azizi.Email)

	anin := User{
		Name:  "Anindhita Cahyadi",
		Age:   24,
		Email: "anin@jkt48.com",
	}

	fmt.Println(anin)
	fmt.Println(anin.Name)
	fmt.Println(anin.Age)
	fmt.Println(anin.Email)

	freya := User{"Freya", "freya@jkt48.com", 18}

	fmt.Println(freya)
	fmt.Println(freya.Name)
	fmt.Println(freya.Age)
	fmt.Println(freya.Email)
}
