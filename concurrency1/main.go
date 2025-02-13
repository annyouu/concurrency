package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text() //読み込んだ文字列をチャネルに送る
		}
		close(ch) //チャネルを閉じる
	}()
	return ch
}

func main() {
	ch := input(os.Stdin)
	for line := range ch {  // チャネルが閉じられるまで繰り返す
		fmt.Println(">", line)
	}
}

// for {
	// 	fmt.Println(">")
	// 	fmt.Println(<-ch)
	// }

// rangeを使ったループ(スライスやマップの反復)
// func main() {
// 	nums := []int{10, 20, 30, 40, 50}
// 	for index, value := range nums {
// 		fmt.Printf("index: %d, value: %d\n", index, value)
// 	}
// }

// マップのループ
// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200, "orange": 300}
// 	for key, value := range m {
// 		fmt.Printf("%s: %d\n", key, value)
// 	}
// }
