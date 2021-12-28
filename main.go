package main

import (
	rand "crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
	"os"
	"time"
	"unsafe"
)

type test struct {
	a int
}

func (t test) fun() {
	t.a++
}

func Float64ToByte(float float64) []byte {
	pointer := *(*[]byte)(unsafe.Pointer(&float))
	return pointer
}

func ByteToFloat64(bytes []byte) float64 {
	return *(*float64)(unsafe.Pointer(&bytes))
}

func chanspeed() {
	n := *flag.Int("n", 100000, "send number times")
	l := *flag.Int("l", 1024000, "bytes size")
	c := make(chan []byte)
	go func(c chan []byte) {
		buf := make([]byte, l)
		for i := 0; i < l; i++ {
			buf[i] = "0123456789abcdef"[i%16]
		}
		begin := time.Now()
		for i := 0; i < n; i++ {
			c <- buf
		}
		since := time.Since(begin).Seconds()
		fmt.Println(since)
		fmt.Printf("speed: %v MB/S", float64(n*l)/1024/1024/since)
		close(c)
	}(c)
	for true {
		_, ok := <-c
		if !ok {
			break
		}
		//fmt.Println(buf)
	}
}
func handlerr(err error) {
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}
}

func rsa_test() {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return
	}
	publicKey := key.PublicKey
	v15, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte("hello"))

	fmt.Println(v15)

	pkcs1v15, err := rsa.DecryptPKCS1v15(rand.Reader, key, v15)

	fmt.Println(string(pkcs1v15))
}

func maximalSquare(matrix [][]byte) int {

	M := make([][]int, len(matrix))
	ans := 0
	for i := 0; i < len(matrix); i++ {
		M[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i > 0 {
					M[i][j] = M[i-1][j] + 1
				} else {
					M[i][j] = 1
				}
				ans = 1
			}
		}
	}
	if ans == 0 {
		return 0
	}
	var stack []int
	stack = append(stack, -1)
	for i := 1; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			for len(stack) > 1 && M[i][stack[len(stack)-1]] >= M[i][j] {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp := min(j-stack[len(stack)-1]-1, M[i][p])
				if tmp > ans {
					ans = tmp
				}
			}
			stack = append(stack, j)
		}
		last := len(M[0])
		for len(stack) > 1 {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := min(last-stack[len(stack)-1]-1, M[i][p])
			if tmp > ans {
				ans = tmp
			}
		}
	}

	return ans * ans
}
func min(args ...int) int {
	ans := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}
func main() {
	print(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}
