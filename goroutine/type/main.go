// 型スイッチは、
// Goでinterface{}型の変数が実際にどの型の値を持っているかを判定し、その型に応じて処理を分岐させる構文のこと

package main

import (
	"fmt"
)

func printType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("整数:", v)
	case string:
		fmt.Println("文字列")
	}
}

func main() {
	printType(100)
	printType("hello")
	printType(true)
	printType(3.14)
}