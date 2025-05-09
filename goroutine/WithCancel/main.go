package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("workerキャンセルを検知しました:", ctx.Err())
		default:
			fmt.Println("worker処理中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	
	time.Sleep(2 * time.Second) // 少し待ってからキャンセルする
	fmt.Println("mainキャンセルを送信")
	cancel()

	time.Sleep(1 * time.Second)
}