package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImgUrls chan string
	waitGroup   sync.WaitGroup
	reImg       = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	chanWork    chan string
)

func main() {
	baseUrl := "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/"
	chanImgUrls = make(chan string, 1000000)
	chanWork = make(chan string, 26)
	for i := 0; i < 26; i++ {
		url := baseUrl + strconv.Itoa(i+1) + ".html"
		waitGroup.Add(1)
		go getImg(url)
	}
	waitGroup.Add(1)
	go checkOk()
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go downImgs()
	}
	waitGroup.Wait()
	fmt.Println("finish downloading img")
}

func checkOk() {
	i := 0
	for s := range chanWork {
		fmt.Println(s + "finished, find the img")
		i++
		if i == 26 {
			break
		}
	}
	fmt.Println("finished all")
	close(chanImgUrls) // close 之后会读完里面的内容然后退出
	close(chanWork)
	waitGroup.Done()
}

func downImgs() {
	for url := range chanImgUrls {
		index := strings.LastIndex(url, "/")
		fileName := url[index+1:]
		fileName = strconv.Itoa(int(time.Now().UnixNano())) + fileName
		ok := downImgUrl(url, fileName)
		if !ok {
			fmt.Printf("%s 下载失败\n", fileName)
		}
	}
	waitGroup.Done()
}

func downImgUrl(url string, fileName string) bool {
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	err = ioutil.WriteFile("./img/"+fileName, body, 0666)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return true
}

func getImg(url string) {
	_, body, err := fasthttp.Get(nil, url)
	handErr(err)
	s := string(body)
	compile := regexp.MustCompile(reImg)
	allString := compile.FindAllString(s, -1)
	fmt.Printf("在 %s 中找到了 %d 个图片\n", url, len(allString))
	for _, s2 := range allString {
		chanImgUrls <- s2
	}
	chanWork <- url
	waitGroup.Done()
}

func handErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
	}
}
