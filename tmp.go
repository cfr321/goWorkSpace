package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
)

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

type result struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (r result) String() string {
	return "{\nmessage:" + r.Message + ",\n" + "status:" + strconv.Itoa(r.Status) + ",\n}"
}

func test1() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}

	log.Println(" ===> loop end.")
}

func health() {
	go http.Get("http://192.168.221.1:12345/health")
	resp, err := http.Get("http://192.168.221.2:12346/health")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(all))
}
func main() {
	health()
}
