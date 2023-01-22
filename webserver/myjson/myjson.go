package myjson

import (
	"encoding/json"
	"fmt"
)

func CreateJSON(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return ""
	}

	return string(jsonData)
}

func CreateStrJSON(data map[string]string) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return ""
	}

	return string(jsonData)
}

func TestJSON() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	fmt.Printf("json data: %s\n", CreateJSON(data))
}

func TestunJSON() {
	jsonData := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`

	type mydata struct {
		IntValue        int
		BoolValue       bool
		StringValue     string
		DateValue       string
		ObjectValue     map[string]interface{}
		NullStringValue *string
		NullIntValue    *int
	}

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	var data2 mydata
	err = json.Unmarshal([]byte(jsonData), &data2)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", data2)
	fmt.Printf("json map: %v\n", data2.DateValue)
	fmt.Printf("json map: %v\n", data2.ObjectValue["arrayValue"])

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

}
