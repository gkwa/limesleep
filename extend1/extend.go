package extend1

import "encoding/json"

type OriginalStruct struct {
	Field1 string
	Field2 int
}

type ExtendedStruct struct {
	OriginalStruct
	ExtraField string
}

func (e *ExtendedStruct) String() string {
	jsonBytes, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return "Error marshaling struct to JSON"
	}
	return string(jsonBytes)
}

func Main() {
	obj := ExtendedStruct{
		OriginalStruct: OriginalStruct{
			Field1: "value1",
			Field2: 42,
		},
		ExtraField: "extraValue",
	}

	println(obj.String())
}
