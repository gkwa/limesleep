package extend2

import (
	"encoding/json"
	"fmt"

	"github.com/taylormonacelli/limesleep/extend1"
)

type extendedStruct struct {
	BaseStruct      extend1.ExtendedStruct
	AdditionalField string
}

func (e *extendedStruct) String() string {
	baseStructBytes, err := json.MarshalIndent(e.BaseStruct, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error marshaling BaseStruct to JSON: %v", err)
	}

	result := fmt.Sprintf("%s\nAdditionalField: %s", string(baseStructBytes), e.AdditionalField)

	return result
}

func Main() {
	obj := extendedStruct{
		BaseStruct: extend1.ExtendedStruct{
			OriginalStruct: extend1.OriginalStruct{
				Field1: "extendedValue1",
				Field2: 84,
			},
			ExtraField: "extendedExtraValue",
		},
		AdditionalField: "additionalValue",
	}

	fmt.Println(obj.String())
	fmt.Println("----------------------------------------")
	fmt.Println(obj.BaseStruct.String())
	fmt.Println("----------------------------------------")
	fmt.Println(obj.BaseStruct)
}
