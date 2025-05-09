package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "dataがついた"
	}()

	select {
	case msg := <- ch:
		fmt.Println("届いた", msg)
	default:
		fmt.Println("dataが届かない")
	}

	// selectの用途
	// select {
	// case res := <- ch:
	// 	fmt.Println("結果:", res)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("タイムアウト")
	// }

	// selectの用途(キャンセル処理)
	// select {
	// case data := <- dataCh:
	// 	fmt.Println("受信:", data)
	// case <-ctx.Done():
	// 	fmt.Println("キャンセルされました:", ctx.Err())
	// }
}