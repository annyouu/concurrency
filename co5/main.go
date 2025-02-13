package main

import (
	"fmt"
	"os"
	"testing"
	"go.uber.org/goleak" //Uberのゴルーチンリーク検出ライブラリ
)

// TestMainは、全てのテストの実行前後に共通処理を行う
func TestMain(m *testing.M) {
	fmt.Println("TestMain: テスト開始") // テスト開始メッセージ

	// goleadを使用してゴルーチンリークを検出
	goleak.VerifyTestMain(m)

	// 全てのテストを実行
	exitCode := m.Run()

	fmt.Println("TestMain: テスト終了") // テスト終了メッセージ
	os.Exit(exitCode) // osに終了コードを通知
}


func TestA(t *testing.T) {
	fmt.Println("TestA: 実行中")
}

func TestB(t *testing.T) {
	fmt.Println("TestB: 実行中")
	t.Fail()
}

// ゴルーチンリークを発生させるテスト
func TestLeaK(t *testing.T) {
	fmt.Println("TestLeak; ゴルーチンリークをさせる")

	go func() {
		select {} // 何もせずに無限に待ち続ける
	}()
}