package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type JsonStr map[string]interface{}

func main() {
	var str JsonStr
	content, _ := ioutil.ReadFile("test.json")
	json.Unmarshal([]byte(content), &str)
	decodeJson(str)
}

var i = 0

func decodeJson(input map[string]interface{}) {
	if i == 0 {
		fmt.Printf("type my_struct_%d struct{\n", i)
	}
	i++
	for k, v := range input {
		//fmt.Println(k)
		switch v := v.(type) {

		default:
			fmt.Printf("%s %s `json:\"%s\"`\n", strings.Title(k), reflect.TypeOf(v), k)

		case map[string]interface{}:
			fmt.Printf("%s struct{\n", strings.Title(k))
			decodeJson(v)
			fmt.Printf(" `json:\"%s\"`\n", k)
//		case []interface{}:
//
//			switch v2 := v[0].(type) {
//			case interface{}:
//				fmt.Printf("%s %s `json:\"%s\"`\n", strings.Title(k), reflect.TypeOf(v), k)
//			case map[string]interface{}:
//				fmt.Printf("%s struct{\n", strings.Title(k))
//				decodeJson(v2)
//				fmt.Printf(" `json:\"%s\"`\n", k)
//			}
//
		}
	}

	fmt.Print("}")

}
