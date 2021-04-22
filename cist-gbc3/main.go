//
// Author: cfr
//

package main

import (

	"fmt"
	Cist "workspace/cist-gbc3/cist"
)

func main() {
	// add test
	var n int
	fmt.Println("请输入n:")
	fmt.Scan(&n)
	Cist.Test(n)
}
