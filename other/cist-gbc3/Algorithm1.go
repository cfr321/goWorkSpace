//
// Author: cfr
//

package cist

type edge struct {
	a, b int
}

type T struct {
	edges []edge
}

//在 K n^2 n^2 的二分图里面构建 n^2/2 个CIST
func BuildCISTinKNN(n int) []T {
	np := n * n
	f := np / 2
	var res []T
	for i := 1; i <= f; i++ {
		var tmp []edge
		tmp = append(tmp, edge{i, i})
		tmp = append(tmp, edge{i, i + f})
		tmp = append(tmp, edge{i + f, i})
		for j := i + 1; j <= i+f-1; j++ {
			tmp = append(tmp, edge{j, i})
			tmp = append(tmp, edge{i, j})
		}
		for j := 1; j <= i-1; j++ {
			tmp = append(tmp, edge{j, i + f})
			tmp = append(tmp, edge{i + f, j})
		}
		for j := i + f + 1; j <= np; j++ {
			tmp = append(tmp, edge{j, i + f})
			tmp = append(tmp, edge{i + f, j})
		}
		res = append(res, T{tmp})
	}
	return res
}
