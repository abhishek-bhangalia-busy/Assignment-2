package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func printJSON(d interface{}, idn string) {
	v := reflect.ValueOf(d)

	switch v.Kind() {
	case reflect.Map:
		fmt.Printf("{ (type : %v, kind : %v)", v.Type(), v.Kind())
		for _, k := range v.MapKeys() {
			fmt.Printf("\n%v\t%v : ", idn, k)
			printJSON(v.MapIndex(k).Interface(), idn+"\t")
		}
		fmt.Printf("\n    %v}", idn)
	case reflect.Slice:
		fmt.Printf("[ (type : %v, kind : %v)", v.Type(), v.Kind())
		for i := 0; i < v.Len(); i++ {
			fmt.Printf("\n%v\t%v : ", idn, i)
			printJSON(v.Index(i).Interface(), idn+"\t")

		}
		fmt.Printf("\n    %v]", idn)
	default:
		fmt.Printf("%v, (type : %v, kind : %v)", v.Interface(), v.Type(), v.Kind())
	}
}

func main() {
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
	var mp interface{}
	err := json.Unmarshal([]byte(inp), &mp)

	if err != nil {
		panic(err)
	}

	printJSON(mp, "")
}
