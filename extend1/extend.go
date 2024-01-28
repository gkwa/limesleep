package extend1

import "fmt"

type originalStruct struct {
	Field1 string
	Field2 int
}

type extendedStruct struct {
	originalStruct
	ExtraField string
}

func Main() {
	obj := extendedStruct{
		originalStruct: originalStruct{
			Field1: "value1",
			Field2: 42,
		},
		ExtraField: "extraValue",
	}

	fmt.Println(obj.Field1)
	fmt.Println(obj.Field2)

	fmt.Println(obj.ExtraField)
}
