// 1.HTTPハンドラ

// 2.HTTPハンドラとエントリポイント（エンドポイント）の結びつけ
// ↑ルーティング

// 3.HTTPサーバの起動

package main

import (
	"fmt"
	"net/http"
)

// HTTPハンドラの作成
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello, world")
}

// HTTPハンドラとエンドポイントの結びつけ
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}