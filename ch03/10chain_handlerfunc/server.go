// リスト3.10
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("ハンドラ関数が呼び出されました - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello2))
	server.ListenAndServe()
}
