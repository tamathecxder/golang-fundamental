package pkg

import (
	"container/list"
	"fmt"
)

func PkgContainerList() {
	data := list.New()

	data.PushBack("Yudistira")
	data.PushBack("Eka")
	data.PushBack("Pratama")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
