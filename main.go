package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var inp string = `{
"name" : "Tolexo Online Pvt. Ltd",
"age_in_years" : 8.5,
"origin" : "Noida",
"head_office" : "Noida, Uttar Pradesh",
"address" : [
{
"street" : "91 Springboard",
"landmark" : "Axis Bank",
"city" : "Noida",
"pincode" : 201301,
"state" : "Uttar Pradesh"
},
{
"street" : "91 Springboard",
"landmark" : "Axis Bank",
"city" : "Noida",
"pincode" : 201301,
"state" : "Uttar Pradesh"
}
],
"sponsers" : {
"name" : "One"
},
"revenue" : "19.8 million$",
"no_of_employee" : 630,
"str_text" : ["one","two"],
"int_text" : [1,3,4]
}`

func getMapItemTypes(mp map[string]interface{}) {
	fmt.Println("\nMap item types :")
	for k, v := range mp {
		rObjTyp := reflect.TypeOf(v)
		rObjVal := reflect.ValueOf(v)
		fmt.Printf("key : %v, value : %v, type : %v, kind : %v\n", k, rObjVal, rObjTyp, rObjVal.Kind())
		checkNestedTypes(rObjTyp, rObjVal, v)
	}
}

func checkNestedTypes(rObjTyp reflect.Type, rObjVal reflect.Value, v interface{}) {
	if rObjTyp.String() != rObjVal.Kind().String() {

		switch rObjVal.Kind().String() {

		case "slice":
			if v, ok := v.([]interface{}); ok {
				getSliceItemTypes(v)
			}

		case "map":
			if v, ok := v.(map[string]interface{}); ok {
				getMapItemTypes(v)
			}

		}
	}
}

func getSliceItemTypes(s []interface{}) {
	fmt.Println("\nSlice Item Types :")
	for _, v := range s {
		rObjTyp := reflect.TypeOf(v)
		rObjVal := reflect.ValueOf(v)
		fmt.Printf("value : %v, type : %v, kind : %v\n", rObjVal, rObjTyp, rObjVal.Kind())
		checkNestedTypes(rObjTyp, rObjVal, v)
	}
}

func main() {
	var mp = make(map[string]interface{})
	err := json.Unmarshal([]byte(inp), &mp)

	if err != nil {
		panic(err)
	}

	getMapItemTypes(mp)
}
