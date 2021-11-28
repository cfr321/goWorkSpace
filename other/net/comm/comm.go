package comm

import "fmt"

func Handerr(err error) {
	if nil != err {
		_ = fmt.Errorf("%v\n", err)
	}
}
