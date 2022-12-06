package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/momeemt/2000s/domain/model"
)

// GetMessageはクライアントに簡単なメッセージを返す関数
func HandleHello(w http.ResponseWriter, r *http.Request) {
	// アクセスを許可したいアクセス元
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//アクセスを許可したいHTTPメソッド
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 許可したいHTTPリクエストヘッダ
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// ヘッダの設定
	r.Header.Set("Content-Type", "application/json")

	// GET
	if r.Method == http.MethodGet {
		// メッセージを作成
		message := model.Message{
			Status:  "200 OK",
			Message: "Hello, World!",
		}

		// JSONに変換
		jsonBytes, err := json.Marshal(message)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, `{"status":"500 INTERNAL SERVER ERROR","message":"JSON Marshal error(Message)"}`)
			fmt.Println("JSON Marshal error(Message)\n", err)
			return
		}

		// Stringに変換
		jsonString := string(jsonBytes)

		// レスポンスを返す
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, jsonString)
	}
}
