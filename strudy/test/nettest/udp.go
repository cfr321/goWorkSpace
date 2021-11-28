package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type easy struct {
	Name string `json:"name"` // 字段解释，可指json 字符串的名字
	Age  int    `json:"age"`
	Like string `json:"like"`
}

func main() {
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Print(request.Method)
		writer.WriteHeader(200)
		json_str, err := json.Marshal(easy{"yxl", 25, "freedom"})
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Print(string(json_str))
		writer.Write(json_str)
	})
	http.ListenAndServe("localhost:8080", nil)
}
