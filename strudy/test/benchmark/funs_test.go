package benchmark

/*
go test -bench=. -benchtime="3s"     运行多少时间
go test -bench=. -benchtime=100000x  运行多少次
go test -bench=BenchmarkSprintf      指定运行的测试函数
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out   打印出更多的东西，以及cpu和内存情况，结合pprof使用
*/
import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10
	//fmt.
	os.is
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
