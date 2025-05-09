package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//レスポンス用の構造体
type HelloResponse struct {
	Message string `json:"message"`
}

// ハンドラー関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//　ステータスコードを設定(200 OK)
	w.WriteHeader(http.StatusOK)

	// JSONでレスポンスを返す
	json.NewEncoder(w).Encode(HelloResponse{Message: "Hello, World!"})
}


func main() {
	// ルーティング設定
	http.HandleFunc("/hello", helloHandler)

	// サーバ起動
	fmt.Println("サーバー起動中 http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバーエラー:", err)
	}
}