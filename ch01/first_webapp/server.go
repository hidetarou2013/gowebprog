package main

// p22
// go run ./ch01/first_webapp/server.go
// go install first_webapp


import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
