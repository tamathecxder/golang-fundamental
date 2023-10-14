package datatypes

import "fmt"

func Maps() {
	oshi := map[string]string{
		"name": "Azizi Shafaa Asadel",
		"age":  "19",
	}

	oshi["height"] = "162"

	fmt.Println(oshi)
	fmt.Println(oshi["name"])
	fmt.Println(oshi["age"])
	fmt.Println(oshi["height"])

	movie := make(map[string]string)
	movie["title"] = "The Grand Budapest Hotel"
	movie["director"] = "Wes Anderson"
	movie["release_date"] = "2014"

	fmt.Println(movie)

	delete(movie, "release_date")

	fmt.Println(movie)
}
