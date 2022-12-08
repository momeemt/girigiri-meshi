package handler

import (
	"fmt"
	"net/http"
)

func ReturnErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintln(w, `{"status":"500 INTERNAL SERVER ERROR","message":"JSON Marshal error(Message)"}`)
	fmt.Println("JSON Marshal error(Message)\n", err)
}
