package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for {
		msg := <- ch // チャネルから受信
		fmt.Printf("Worker %d: %s\n", id, msg)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker(1, ch1)
	go worker(2, ch2)

	go func() {
		for {
			ch1 <- "Hello from ch1"
			time.Sleep(1 * time.Second)
		}
	}()
	
	go func() {
		for {
			ch2 <- "Hello from ch2"
			time.Sleep(2 * time.Second)
		}
	}()

	// メインのゴルーチンはselectで受信
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Main received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Main received:", msg2)
		}
	}
}