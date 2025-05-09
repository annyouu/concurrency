// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <- ctx.Done():
// 			fmt.Println("タイムアウトorキャンセル検知:", ctx.Err())
// 			return
// 		default:
// 			fmt.Println("処理中...")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
// 	defer cancel()

// 	go worker(ctx)

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("mainの終了")
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("処理が完了しました")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // リソース解放のため呼ぶ

	err := doSomething(ctx)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("正常終了")
	}
}