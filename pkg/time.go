package pkg

import (
	"fmt"
	"time"
)

func PkgTime() {
	birthDate := "2023-01-12"

	today := time.Now()
	// date := time.Date()

	fmt.Println(today.Year())
	fmt.Println(today.Month())
	fmt.Println(today.Day())
	fmt.Println(today.Hour())
	fmt.Println(today.Minute())
	fmt.Println(today.Second())
	fmt.Println(today.Nanosecond())

	utc := time.Date(2023, 10, 23, 11, 42, 10, 10, time.UTC)

	fmt.Println("UTC:", utc)

	layout := "2006-01-02"
	formattedDate, _ := time.Parse(layout, birthDate)

	fmt.Println(formattedDate)

}
