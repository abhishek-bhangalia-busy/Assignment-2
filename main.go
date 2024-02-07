package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
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

func printSliceItems(s []interface{}) {
	for i, v := range s {

		rObjVal := reflect.ValueOf(v)
		checkNestedTypes(rObjVal, strconv.Itoa(i))
	}
}

func checkNestedTypes(rObjVal reflect.Value, k string) {
	switch rObjVal.Kind() {
	case reflect.Map:
		printJSON(rObjVal.Interface().(map[string]interface{}))
	case reflect.Slice:
		printSliceItems(rObjVal.Interface().([]interface{}))
	default:
		rObjTyp := reflect.TypeOf(rObjVal.Interface())
		fmt.Printf("key : %v, value : %v, type : %v, kind : %v\n", k, rObjVal, rObjTyp, rObjVal.Kind())
	}
}

func printJSON(mp map[string]interface{}) {
	for k, v := range mp {
		rObjVal := reflect.ValueOf(v)
		checkNestedTypes(rObjVal, k)
	}
}

func main() {
	var mp = make(map[string]interface{})
	err := json.Unmarshal([]byte(inp), &mp)

	if err != nil {
		panic(err)
	}

	printJSON(mp)
}
