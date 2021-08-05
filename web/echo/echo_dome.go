package main

import (
	"log"
	"net/http"
)

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	println(request.RequestURI)
	_, _ = writer.Write([]byte("hello"))
}

func main() {
	err := http.ListenAndServe(":8080", MyHandler{})
	if err != nil {
		log.Fatal(err)
	}
}
