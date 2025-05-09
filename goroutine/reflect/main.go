package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	u := User{"Taro", 30}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %s, JSON Tag: %s\n", field.Name, tag)
	}
}