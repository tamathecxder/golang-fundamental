package pkg

import (
	"fmt"
	"regexp"
)

func PkgRegexp() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])a")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eka"))
	fmt.Println(regex.MatchString("eva"))

	search := regex.FindAllString("ema eqa elo eza", -1)

	fmt.Println(search)
}
