package functions

import "fmt"

func askFavoriteSong(artist string, title string) (string, string) {
	return artist, title
}

func MultipleReturns() {
	band1, title1 := askFavoriteSong("The Beatles", "Here, There and Everywhere")

	fmt.Println(band1, title1)

	// ignoring return value
	singer1, _ := askFavoriteSong("Phoebe Bridgers", "")

	fmt.Println(singer1)
}
