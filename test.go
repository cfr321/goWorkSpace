package main

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strings"
	"workspace/myUtil"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func trap(height []int) int {
	var s []int
	res := 0
	//s = append(s, 0)
	for i := 0; i < len(height); i++ {
		for len(s) > 0 && height[i] >= height[s[len(s)-1]] {
			t := s[len(s)-1]
			s = s[0 : len(s)-1]
			if len(s) > 0 {
				before := s[len(s)-1]
				res += (min(height[i], height[before]) - height[t]) * (i - before)
			}
		}
		s = append(s, i)
	}
	return res
}

func trap1(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for left < right {
		leftMax = int(math.Max(float64(leftMax), float64(height[left])))
		rightMax = int(math.Max(float64(rightMax), float64(height[right])))
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

type Node struct {
	a, b int
}

func (n Node) Less(other myUtil.CmpInterface) bool {
	return n.b < other.(Node).b
}

type easy struct {
	Name string `json:"name"` // 字段解释，可指json 字符串的名字
	Age  int    `json: age`
	Like string `json: like`
}

func handle(writer http.ResponseWriter, request *http.Request) {
	data := []byte("helloNihao")
	writer.Write(data)
}

func handle2(writer http.ResponseWriter, request *http.Request) {
	data := []byte("草拟吗")
	writer.Write(data)
}

type student struct {
	name string
	size int
	sex  string
}

func (s student) Read(p []byte) (n int, err error) {
	return 0, err
}

type Person struct {
	name string
}

func storeWater(bucket []int, vat []int) int {
	ans := math.MaxInt32
	for i := 1; i < 10000; i++ {
		cur := 0
		for j := 0; j < len(vat); j++ {
			k := (vat[j] + i - 1) / i
			cur += max(0, k-bucket[j])
		}
		ans = min(ans, i+cur)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	app := false
	left := newInterval[0]
	right := newInterval[1]
	for _, interval := range intervals {
		if interval[0] > right {
			if !app {
				res = append(res, []int{left, right})
				app = true
			}
			res = append(res, interval)
		} else if interval[1] < left {
			res = append(res, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !app {
		res = append(res, []int{left, right})
	}
	return res
}

var ans1723 int

type work struct {
	id       int
	workTime int
}

var works []work

func minimumTimeRequired(jobs []int, k int) int {
	works = make([]work, k)
	for i := 0; i < k; i++ {
		works[i].workTime = 0
		works[i].id = i
	}
	ans1723 = math.MaxInt32
	dfs1723(0, jobs, 0)
	return ans1723
}

func dfs1723(i int, jobs []int, tmp int) {
	if i == len(jobs) {
		if tmp < ans1723 {
			ans1723 = tmp
		}
		return
	}
	sort.Slice(works, func(i, j int) bool {
		return works[i].workTime < works[j].workTime
	})
	cop := make([]work, len(works))
	copy(cop, works)
	for j := 0; j < len(works); j++ {
		if j == 0 || works[j].workTime != works[j-1].workTime {
			if works[j].workTime+jobs[i] < ans1723 {
				works[j].workTime += jobs[i]
				tmpId := works[j].id
				dfs1723(i+1, jobs, max(tmp, works[j].workTime))
				for k := 0; k < len(works); k++ {
					if works[k].id == tmpId {
						works[k].workTime -= jobs[i]
						break
					}
				}
			}
		}
	}
}

var c = flag.Bool("c", false, "something")

type Noname interface {
	add()
	// a string
}

type Sun struct {
	a int
}

func (s Sun) add() {
	s.a++
}

func Newnoname() Noname {
	return &Sun{1}
}
func add(s Sun) {

}
func main() {
	s := &Sun{1}
	s.add()
}


func readDir(path string) {
	dir, _ := ioutil.ReadDir(path)
	for _, file := range dir {
		filePath := path + "/" + file.Name()
		if !file.IsDir() {
			pocessFile(filePath)
		} else {
			readDir(filePath)
		}
	}
}
func pocessFile(filePath string) {
	if strings.HasSuffix(filePath, ".java") {
		content, _ := ioutil.ReadFile(filePath)
		all := strings.ReplaceAll(string(content), "杨德石", "Lyy")
		_ = ioutil.WriteFile(filePath, []byte(all), fs.ModeAppend)
	}
}

func funcName() {
	ints := make([]int, 30000000)
	for i := 0; i < 30000000; i++ {
		ints[i] = 1
	}
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

