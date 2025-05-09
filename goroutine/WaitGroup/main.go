package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 処理が終わったことを通知する
	fmt.Printf("Worker %d is working...\n", id)
}

func main() {
	var wg sync.WaitGroup

	// 3つのゴルーチンを待つ
	for i := 1; i <= 3; i++ {
		wg.Add(1) // ゴルーチンを起動する前にAddする
		go worker(i, &wg)
	}

	wg.Wait() // 全てのDone()を待つ
	fmt.Println("全てのゴルーチンが終了")

	// 間違い
	// 1 ゴルーチン内でAdd()を呼んでしまう
	// for i := 0; i < 3; i++ {
	// 	go func() {
	// 		wg.Add(1) //　ゴルーチン内で呼ぶと競合が起きやすい
	// 		defer wg.Done()
	// 	}()
	// }
	// wg.Wait()  // Wait()がAdd()より先に呼ばれると、カウントが0のままなので即座に終了する

	// 2 Done()の呼び忘れ：Done()を呼ばなければ、Wait()は永遠に終わらない(デットロック)

	// 3 WaitGroupを値渡しする
	// (WaitGroupはポインタ渡し(*sync.WaitGroup)しないとAddとDoneが異なる変数を使ってしまう)

	// 4 Add()とDone()をゴルーチン内で実行する
	// WaitはAdd()が完了した後に呼ぶべき
}


