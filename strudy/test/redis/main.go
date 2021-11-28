package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RunAble func(int, int)
type Ex struct {
	RunAble
}

func (e Ex) Process(a int, b int) {
	fmt.Println(a + b)
}
func (a RunAble) name() {
	a(1, 2)
}

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.29.130:6379",
		Password: "",
		DB:       0,
	})

	result, _ := client.Do(ctx, "get", "hello").Result()
	fmt.Println(result)
}
