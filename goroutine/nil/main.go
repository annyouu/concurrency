package main

import (
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string { 
	return "error" 
}

func test() error {
	var e *MyError = nil // eはMyError型のポインタで、値はnilという意味。
	return e
}

func main() {
	err := test()
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is not nil")
	}
}