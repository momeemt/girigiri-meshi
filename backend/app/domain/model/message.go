package model

// Message は実際にクライアントに返信するデータの構造体
type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
