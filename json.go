package main

import (
    "encoding/json"
    "net/http"
)

// jsonタグを使ってjsonのkeyを指定できる
type user struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Birth string `json:"birth"`
}

func makeJson(w http.ResponseWriter, r *http.Request) {
    userData := &user{
        Name:  "takumaron",
        Age:   20,
        Birth: "2020/05/03",
    }

		// HTTPヘッダーの追加（json形式のデータであることと、utf-8の文字列であること）
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		// HTTPステータスコードの追加（200,ステータスOK）
		w.WriteHeader(http.StatusOK)
		
		// userDataをjson形式に変換
    data, err := json.Marshal(userData)
    if err != nil {
        http.Error(w, err.Error(), 500)
		}
		// json化したWriteデータの送信
    _, _ = w.Write(data)
}

func main() {
    http.HandleFunc("/", makeJson)
    http.ListenAndServe(":8080", nil)
}