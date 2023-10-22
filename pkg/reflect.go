package pkg

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return false
			}
		}
	}

	return true
}

func PkgReflect() {
	sample := Sample{"Function"}

	sampleType := reflect.TypeOf(sample)
	sampleNumField := sampleType.NumField()
	sampleName := sampleType.Field(0).Name
	sampleField := sampleType.Field(0)
	sampleTag := sampleField.Tag.Get("required")

	fmt.Println(sampleType)
	fmt.Println(sampleNumField)
	fmt.Println(sampleName)
	fmt.Println(sampleTag)

	fmt.Println("Is Valid?", isValid(sample))
}
