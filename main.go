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

func main() {
	fmt.Println("hello")
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
