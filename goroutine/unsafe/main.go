package main

import (
	"fmt"
)

type Bar interface {
	DoSomething()
}

type Foo struct {}

func (f *Foo) DoSomething() {
	fmt.Println("doing something")
}

// インターフェース実装チェック
var _ Bar = (*Foo)(nil)