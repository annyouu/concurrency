package main

import(
	"fmt"
)

// チャネルを作成して返す
func makeCh() chan int {
	return make(chan int)
}

// チャネルから値を受信して返す
func recvCh(recv chan int) int {
	return <- recv
}

func main() {
	ch := makeCh()
	go func() {
		ch <- 100
	}()
	fmt.Println(recvCh(ch))
}


// func newChannel() chan string {
// 	return make(chan string)
// }

// func main() {
// 	ch := newChannel()
// 	go func() {
// 		ch <- "Hello"
// 	}()
// 	fmt.Println(<-ch)  //関数の戻り値として<-chを渡す
// }


// func sendData(ch chan int) {
// 	ch <- 100
// }

// func main() {
// 	ch := make(chan int)   //関数の引数に渡す
// 	go sendData(ch)
// 	fmt.Println(<-ch)
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan string)
// 	go func() {
// 		ch1 <- 100
// 	}()

// 	go func() {
// 		ch2 <- "hi"
// 	}()

// 	select {
// 	case v1 := <-ch1:
// 		fmt.Println(v1)
	
// 	case v2 := <-ch2:
// 		fmt.Println(v2)
// 	}

// }