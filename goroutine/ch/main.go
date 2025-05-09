// 終了検知するrange ch
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// 受信側
	go func() {
		// msg := <-ch  受信
		for v := range ch {
			fmt.Println("受信:", v)
		}
		fmt.Println("チャネルがcloseされました")
	}()

	// 送信側
	for i := 0; i < 3; i++ {
		ch <- i
	}

	close(ch)
}
