package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)


// 一回だけinitし、初期化する
func init() {
	rand.Seed(time.Now().UnixNano())
}

func omikuji(w http.ResponseWriter, r *http.Request)  {
	switch rand.Intn(6) {
	case 0:
		fmt.Fprint(w, "凶")
	case 1, 2:
		fmt.Fprint(w, "小吉")
	case 3, 4:
		fmt.Fprint(w, "中吉")
	case 5:
		fmt.Fprint(w, "大吉")
	}
}

// HTTPハンドラとエンドポイントの結びつけ
func main() {
	http.HandleFunc("/", omikuji)
	http.ListenAndServe(":8080", nil)
}