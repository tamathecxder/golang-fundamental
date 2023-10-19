package _import

import (
	"fmt"
	"golang-fundamental/access-modifier/_export"
)

func Consume() {
	fmt.Println(_export.AppName)
	fmt.Println(_export.VersionInfo())
}
