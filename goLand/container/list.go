//
// Author: cfr
//

package container

import (
	"container/list"
	"container/ring"
)

func Listtest() {
	l:= list.New()
	l.PushBack(1)
	r := ring.New(5)
}