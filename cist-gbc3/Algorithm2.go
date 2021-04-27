//
// Author: cfr
//

package cist

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"syscall"
)

var globeN int
var np int
var t int

//节点定义，三个下标定位一个点
type node struct {
	h, i, j int
}

//为方便，重命名[]node，表示一个点集
type nodeArray []node

//边
type Edge struct {
	left, right node
}

//为方便，重命名[]Edge，表示一个边集
type edgeArray []Edge

//也是边集合，暂时只有一个边的数组，后面可以加一些属性，处理边集合
//
type Ek struct {
	edges []Edge
}

//这个可以表示一个BCubeN1在各个CIST中内节点集的集合
type InnerNodeArraysInOneBCubeN1 struct {
	allOnlyOne bool
	maxLenId   int
	innerNodes []nodeArray
}

// 对 主 内节点"排序",保证内节点 前 n 个里面包含 一个 Bcube（n,0）循环，也就是要
func (this *InnerNodeArraysInOneBCubeN1) sort(n int) {
	rem := make([]bool, n)
	num := 0
	array := this.innerNodes[this.maxLenId]
	for i, n2 := range array {
		if !rem[n2.j%n] {
			rem[n2.j%n] = true
			array[num], array[i] = array[i], array[num]
			num++
			if num == n {
				break
			}
		}
	}
}

func (this InnerNodeArraysInOneBCubeN1) minLenId() int {
	res := 0
	for i := 1; i < len(this.innerNodes); i++ {
		if len(this.innerNodes[i]) < len(this.innerNodes[res]) {
			res = i
		}
	}
	return res
}
func (this *InnerNodeArraysInOneBCubeN1) computerLen() {
	this.allOnlyOne = true
	this.maxLenId = 0
	for i := 0; i < len(this.innerNodes); i++ {
		if len(this.innerNodes[i]) > 1 {
			this.allOnlyOne = false
		}
		if len(this.innerNodes[i]) > len(this.innerNodes[this.maxLenId]) {
			this.maxLenId = i
		}
	}
}

//算法二
func BuildCISTsInLGBCN31(n int, tTrees []T) []Ek {
	globeN = n
	t = (n + 4) / 4 // t = (n+1)/4 向上取整
	np = n * n
	var res []Ek

	res = make([]Ek, t)

	Vup := make([]InnerNodeArraysInOneBCubeN1, np) //一个VsInBCubeN1，就是一个bcube的k个独立生成树下的内节点
	Vdown := make([]InnerNodeArraysInOneBCubeN1, np)
	for i := 0; i < np; i++ {
		Vup[i].innerNodes = make([]nodeArray, t)
		Vdown[i].innerNodes = make([]nodeArray, t)
	}

	for k := 0; k < t; k++ {
		var Ek Ek
		var Vk0 []nodeArray //第 k 课独立生成树，上面的所有Bcube(n,1)的内节点
		var Vk1 []nodeArray //第 k 课独立生成树，下面的所有Bcube(n,1)的内节点
		for i := 0; i < np; i++ {
			Vk0 = append(Vk0, nodeArray{})
			Vk1 = append(Vk1, nodeArray{})
		}
		for _, e := range tTrees[k].edges {
			nodea := node{0, e.a - 1, e.b - 1}
			nodeb := node{1, e.b - 1, e.a - 1}
			Vk0[e.a-1] = append(Vk0[e.a-1], nodea)
			Vk1[e.b-1] = append(Vk1[e.b-1], nodeb)
			Ek.edges = append(Ek.edges, Edge{nodea, nodeb})
		}
		//fmt.Printf("第%d课独立生成树跨Bcube边，%v\n\n",k+1,Ek.edges)
		res[k].edges = append(res[k].edges, Ek.edges...)

		for index, array := range Vk0 {
			Vup[index].innerNodes[k] = array
		}
		for index, array := range Vk1 {
			Vdown[index].innerNodes[k] = array
		}
	}

	for i := 0; i < np; i++ {
		// h=0
		W := getW(Vup[i])
		Vup[i].computerLen() //计算各个内节点集合长度，获取长度是否都为一，以及最长的那个下标
		if Vup[i].allOnlyOne {
			cisTsInBcubei := LNode_CISTs(Vup[i])
			for y := 0; y < t; y++ {
				res[y].edges = append(res[y].edges, cisTsInBcubei[y]...) //把对应的 Bcube的第 i 课独立生成树的边加到 全局 的第 i 课独立生成树
			}
		} else {
			cisTsInBcubei := INode_CISTs(Vup[i], W)
			for y := 0; y < t; y++ {
				res[y].edges = append(res[y].edges, cisTsInBcubei[y]...)
			}
		}
	}

	for i := 0; i < np; i++ {
		// h =1
		W := getW(Vdown[i])
		Vdown[i].computerLen()
		if Vdown[i].allOnlyOne {
			cisTsInBcubei := LNode_CISTs(Vdown[i])
			for y := 0; y < t; y++ {
				res[y].edges = append(res[y].edges, cisTsInBcubei[y]...) //把对应的 Bcube的第 i 课独立生成树的边加到 全局 的第 i 课独立生成树
			}
		} else {
			cisTsInBcubei := INode_CISTs(Vdown[i], W)
			for y := 0; y < t; y++ {
				res[y].edges = append(res[y].edges, cisTsInBcubei[y]...)
			}
		}
	}
	return res
}

/*
help to build W
*/
func getW(Vs InnerNodeArraysInOneBCubeN1) []bool {
	rem := make([]bool, np) //用来记录出现过的节点
	// build W
	for _, innerNodes := range Vs.innerNodes {
		for _, nd := range innerNodes {
			rem[nd.j] = true
		}
	}
	for i := 0; i < np; i++ {
		rem[i] = !rem[i]
	}
	return rem
}

/*
Case1
*/
func LNode_CISTs(Vis InnerNodeArraysInOneBCubeN1) []edgeArray {

	h := Vis.innerNodes[0][0].h //从中取一个点来获取当前的 h 和 i
	i := Vis.innerNodes[0][0].i

	var Eits []edgeArray //第i课Bcude的独立生成树集合
	n := globeN
	alph := n / 2
	for k := 0; k < t; k++ {
		r := Vis.innerNodes[k][0].j
		var beida int
		if r%n+alph < n {
			beida = r + alph
		} else {
			beida = r + alph - n
		}
		// Vis.innerNodes[k] = append(Vis.innerNodes[k], node{h, i, beida})

		// 点其实不需要处理，能求出 a b即可
		// r%=n
		// beida%=n
		r %= n
		beida %= n
		var a, b int
		if r < beida {
			a = r
			b = beida
		} else {
			a = beida
			b = r
		}

		var Eik edgeArray
		for j := 0; j < n; j++ {
			for m := a + 1; m <= b; m++ {
				node1 := node{h, i, j*n + a}
				node2 := node{h, i, j*n + m}
				Eik = append(Eik, Edge{node1, node2})
			}
			for m := b + 1; m < n; m++ {
				node1 := node{h, i, j*n + b}
				node2 := node{h, i, j*n + m}
				Eik = append(Eik, Edge{node1, node2})
			}
			for m := 0; m < a; m++ {
				node1 := node{h, i, j*n + b}
				node2 := node{h, i, j*n + m}
				Eik = append(Eik, Edge{node1, node2})
			}
			if j < n-1 {
				node1 := node{h, i, b}
				node2 := node{h, i, (j+1)*n + b}
				Eik = append(Eik, Edge{node1, node2})
			}
		}
		Eits = append(Eits, Eik)
	}
	return Eits
}

/*
Case2
*/
func INode_CISTs(Vis InnerNodeArraysInOneBCubeN1, W []bool) []edgeArray {

	h := Vis.innerNodes[0][0].h //从中取一个点来获取当前的 h 和 i
	i := Vis.innerNodes[0][0].i

	sigma := Vis.maxLenId // 获取最多内节点的那个独立生成树

	Vis.sort(globeN)

	//if h == 0 && i == 13 {
	//	fmt.Println(Vis.innerNodes[sigma])
	//}

	n := globeN
	// 将点分入不同的CIST当内节点
	distributeNode(&Vis, W, sigma, n, h, i)
	// 到此节点分类已经完成

	//if h == 0 && i == 3 {
	//	for _, innerNode := range Vis.innerNodes {
	//		fmt.Println("=======")
	//		for _, n2 := range innerNode {
	//			fmt.Print(n2.j)
	//			fmt.Print(",")
	//		}
	//	}
	//}

	//var Eits []edgeArray //第i课Bcude的独立生成树集合
	Eits := make([]edgeArray, t)

	//将各个独立生成树内节点先连起来

	for k := 0; k < t; k++ {
		Vik := Vis.innerNodes[k] //点集合
		var Eik edgeArray
		Eik = connectInerNodes(Vik) // 在 Bcube(n,1)里将内节点链接起来

		Eits[k] = append(Eits[k], Eik...)
	}

	// 完成 sigma 那颗树的生成
	//if i < t {
	//	EiOfSigma := FPath(sigma, Vis)
	//	Eits[sigma] = append(Eits[sigma], EiOfSigma...) //把这些点也加到对应的生成树里面
	//} else {
	//	EiOfSigma := LPath(sigma, Vis)
	//	Eits[sigma] = append(Eits[sigma], EiOfSigma...) //把这些点也加到对应的生成树里面
	//}
	//
	//path := VPath(sigma, Vis)
	//for i := 0; i < t; i++ {
	//	if i != sigma {
	//		Eits[i] = append(Eits[i], path[i]...)
	//	}
	//}
	return Eits //返回第i课Bcube(n,1)的所有独立生成树边
}

var BcuMap [][]int
var visit []bool
var res edgeArray

func connectInerNodes(vik nodeArray) edgeArray {
	res = res[:0]
	visit = make([]bool, np)
	for i := range visit {
		visit[i] = true
	}
	for _, node := range vik {
		visit[node.j] = false
	}
	dfs(vik[0].j, vik[0].h, vik[0].i)
	for i := 0; i < len(visit); i++ {
		if !visit[i] {
			fmt.Printf("dis connect %d  %d\n", vik[0].h, vik[0].i)
		}
	}
	return res
}
func dfs(i int, h, t int) {
	visit[i] = true
	for j, con := range BcuMap[i] {
		if con == 1 && !visit[j] {
			res = append(res, Edge{node{h, t, i}, node{h, t, j}})
			dfs(j, h, t)
		}
	}
}

//创建Bcube(n,1) => BcuMap
func buildBcube(n int) {
	npp := n * n
	BcuMap = BcuMap[:0]
	for i := 0; i < npp; i++ {
		tmp := make([]int, npp)
		BcuMap = append(BcuMap, tmp)
	}
	for i := 0; i < npp; i++ {
		k := i / n
		for j := k * n; j < k*n+n; j++ {
			BcuMap[i][j] = 1
		}
		for j := i % n; j < npp; j += n {
			BcuMap[i][j] = 1
		}
	}
}

// 节点分入不同的独立生成树
func distributeNode(Vis *InnerNodeArraysInOneBCubeN1, W []bool, sigma int, n int, h int, i int) {
	var rk, p, time int
	for round := 0; round < 2; round++ {
		for k := 0; k < t; k++ {
			if k != sigma {
				Vik := &Vis.innerNodes[k] //便于操作
				rk = (*Vik)[0].j
				if round == 0 {
					p = 0
				} else {
					p = 1
					time = 0
				}
				var c = -rk / n
				for p != n {
					if time == 2 && round == 1 {
						p++
					}
					pos := rk + p + c*n
					if pos < np && W[pos] {
						*Vik = append(*Vik, node{h, i, pos})
						W[pos] = false
						p++
						c++
						time = 0
					} else {
						c++
					}
					if rk+p+c*n >= np || rk+p+c*n < 0 {
						c = -rk / n
						time++
					}
				}
			}
		}
	}

	//将剩余的点加到以此加入
	for i2 := 0; i2 < np; i2++ {
		if W[i2] { //代表点还在剩余集合中
			id := Vis.minLenId() //获取内联点集集合最短的
			Vis.innerNodes[id] = append(Vis.innerNodes[id], node{h, i, i2})
		}
	}
	if h == 0 && i == 2 {
		for i1, innerNode := range Vis.innerNodes {
			fmt.Printf("let node%d = [", i1+1)
			for _, n2 := range innerNode {
				fmt.Printf("%d,", n2.j)
			}
			fmt.Print("]")
			fmt.Println()
		}
	}
}

/*
	构造sigma那课 CIST   i<t
*/
func FPath(sigma int, Vis InnerNodeArraysInOneBCubeN1) edgeArray {
	h := Vis.innerNodes[0][0].h //从中取一个点来获取当前是那个一Bcube(n,1)，及得到 h 和 i
	i := Vis.innerNodes[0][0].i

	Vsigma := Vis.innerNodes[sigma] // 取出 sigma 那颗独立生成树的内节点集合

	var Eisigma edgeArray
	rsigma := Vsigma[len(Vsigma)-1].j

	n := globeN

	dT := make([]bool, np) //记录节点是不是被链接过
	//链接 r-n-p 和 一个 r -p +xn
	for k := 0; k < t; k++ {
		if k != sigma {

			//用来记录 Vis[k]中有哪些节点,但是不要前n个
			rem := make([]bool, np)
			for i := n; i < len(Vis.innerNodes[k]); i++ {
				rem[Vis.innerNodes[k][i].j] = true
			}

			for p := 1; p < n; p++ {
				//首先需要估计 x 的取值范围
				// 0 <= r-p+xn < n^2  ==>     (p-r)/n <= x < (n^2-r+p)/n
				for x := (p-rsigma)/n - 1; x < n+1; x++ { //稍稍扩大，边界太难扣了。。。
					pos := rsigma - p + x*n
					if pos < 0 {
						continue
					}
					if pos >= np {
						break
					}
					//找一个
					if rem[pos] {
						Eisigma = append(Eisigma, Edge{node{h, i, rsigma - n - p}, node{h, i, pos}})
						dT[pos] = true
						break
					}
				}

			}
		}
	}
	//链接 r-p  和 剩余  r-p+xn
	for k := 0; k < t; k++ {
		if k != sigma {
			//用来记录 Vis[k]中有哪些节点，这里全部记下来
			rem := make([]bool, np)
			for _, n2 := range Vis.innerNodes[k] {
				rem[n2.j] = true
			}

			for p := 0; p < n; p++ {
				//首先需要估计 x 的取值范围
				// 0 <= r-p+xn < n^2  ==>     (p-r)/n <= x < (n^2-r+p)/n
				for x := (p-rsigma)/n - 1; x < n+1; x++ { //稍稍扩大，边界太难扣了。。。
					pos := rsigma - p + x*n
					if pos < 0 {
						continue
					}
					if pos >= np {
						break
					}
					//找多个
					if rem[pos] && !dT[pos] {
						Eisigma = append(Eisigma, Edge{node{h, i, rsigma - p}, node{h, i, pos}})
						dT[pos] = true
						//break  这里不break，找所有
					}
				}
			}
		}
	}
	return Eisigma
}

/*
	构造sigma那课 CIST   i>=t
*/
func LPath(sigma int, Vis InnerNodeArraysInOneBCubeN1) edgeArray {
	h := Vis.innerNodes[0][0].h //从中取一个点来获取当前是那个一Bcube(n,1)，及得到 h 和 i
	i := Vis.innerNodes[0][0].i
	n := globeN

	Vsigma := Vis.innerNodes[sigma] //sigma这颗树内节点集合

	// rsigma := Vsigma[0].j   //算法没有用到这个

	var Eisigma edgeArray //结果边集合

	dT := make([]bool, np) //记录节点是不是被链接过

	for k := 0; k < t; k++ {
		if k != sigma {

			//用来记录 Vis[k]中有哪些节点,但是不要前n个
			rem := make([]bool, np)
			for i := n; i < len(Vis.innerNodes[k]); i++ {
				rem[Vis.innerNodes[k][i].j] = true
			}
			for p := 0; p < n; p++ {
				if p != sigma {
					a := Vsigma[p+n].j //这这
					for x := -a/n - 1; x < n; x++ {
						pos := a + x*n
						if pos < 0 {
							continue
						}
						if pos >= np {
							break
						}
						if rem[pos] && !dT[pos] {
							Eisigma = append(Eisigma, Edge{node{h, i, a}, node{h, i, pos}})
							dT[pos] = true
							break
						}
					}

				}
			}
		}
	}

	for k := 0; k < t; k++ {
		if k != sigma {

			//用来记录 Vis[k]中有哪些节点,但是要前n个
			rem := make([]bool, np)
			for i := 0; i < len(Vis.innerNodes[k]); i++ {
				rem[Vis.innerNodes[k][i].j] = true
			}
			for p := 0; p < n; p++ {
				a := Vsigma[p].j
				for x := -a/n - 1; x < n; x++ {
					pos := a + x*n
					if pos < 0 {
						continue
					}
					if pos >= np {
						break
					}
					if rem[pos] && !dT[pos] {
						Eisigma = append(Eisigma, Edge{node{h, i, a}, node{h, i, pos}})
					}
				}
			}
		}
	}

	return Eisigma
}

func deleteNode(array *nodeArray, pos int) {
	for i, n := range *array {
		if n.j == pos {
			*array = append((*array)[0:i], (*array)[i+1:]...)
			break
		}
	}
}

/*
	构造其余  CIST
*/
func VPath(sigma int, Vis InnerNodeArraysInOneBCubeN1) []edgeArray {
	h := Vis.innerNodes[0][0].h //从中取一个点来获取当前是那个一Bcube(n,1)，及得到 h 和 i
	i := Vis.innerNodes[0][0].i
	n := globeN

	Eks := make([]edgeArray, t)

	distributeVks := make([][]nodeArray, t)
	for index, innerNode := range Vis.innerNodes {
		distributeVks[index] = distributeVk(innerNode)
	}

	for k := 0; k < t; k++ {
		Vsigma := make(nodeArray, len(Vis.innerNodes[sigma]))
		copy(Vsigma, Vis.innerNodes[sigma]) // 取出 sigma 那颗独立生成树的内节点集合
		var rsigma int
		if i < t {
			rsigma = Vsigma[len(Vsigma)-1].j //the last vertex in Bcube(n,1)
		} else {
			rsigma = Vsigma[sigma].j
		}
		dT := make([]bool, np) //用来记录 k 的内节点是否有被连结果
		if k != sigma {
			if i < t {
				//用来记录 Vis[k]中有哪些节点,但是不要前n个
				rem := make([]bool, np)
				for i := n; i < len(Vis.innerNodes[k]); i++ {
					rem[Vis.innerNodes[k][i].j] = true
				}
				for p := 1; p < n; p++ {
					//首先需要估计 x 的取值范围
					// 0 <= r-p+xn < n^2  ==>     (p-r)/n <= x < (n^2-r+p)/n
					for x := (p-rsigma)/n - 1; x < n+1; x++ { //稍稍扩大，边界太难扣了。。。
						pos := rsigma - p + x*n
						if pos < 0 {
							continue
						}
						if pos >= np {
							break
						}
						//找一个
						if rem[pos] {
							Eks[k] = append(Eks[k], Edge{node{h, i, rsigma - p}, node{h, i, pos}})
							dT[pos] = true
							dT[rsigma-p] = true
							deleteNode(&Vsigma, rsigma-p)
							break
						}
					}
				}
			} else {
				//用来记录 Vis[k]中有哪些节点,但是不要前n个
				rem := make([]bool, np)
				for i := n; i < len(Vis.innerNodes[k]); i++ {
					rem[Vis.innerNodes[k][i].j] = true
				}
				for p := 0; p < n; p++ {
					if p != sigma {
						a := Vis.innerNodes[sigma][p].j
						for x := -a/n - 1; x <= n; x++ {
							pos := a + x*n
							if pos < 0 {
								continue
							}
							if pos >= np {
								break
							}
							if rem[pos] {
								Eks[k] = append(Eks[k], Edge{node{h, i, a}, node{h, i, pos}})
								dT[pos] = true
								dT[a] = true
								deleteNode(&Vsigma, a)
								break
							}
						}
					}
				}
			}
			deleteNode(&Vsigma, rsigma)
			rem := make([]bool, np)
			for i := 0; i < len(Vsigma); i++ {
				rem[Vsigma[i].j] = true
			}

			for len(Vsigma) > 0 {
				for s := 0; s < n; s++ {
					b := Vis.innerNodes[k][s].j
					for x := -b/n - 1; x <= n; x++ {
						pos := b + x*n
						if pos < 0 {
							continue
						}
						if pos >= np {
							break
						}
						if rem[pos] {
							Eks[k] = append(Eks[k], Edge{node{h, i, b}, node{h, i, pos}})
							rem[pos] = false
							dT[b] = true
							dT[pos] = true
							deleteNode(&Vsigma, pos)
						}
					}
				}
			}

			y := rsigma / n

			//记录下 Vis[k]中有的节点
			rem = make([]bool, np)
			for i := 0; i < len(Vis.innerNodes[k]); i++ {
				rem[Vis.innerNodes[k][i].j] = true
			}
			for d := 0; d < n; d++ {
				if rem[y*n+d] {
					Eks[k] = append(Eks[k], Edge{node{h, i, rsigma}, node{h, i, y*n + d}})
					dT[y*n+d] = true
					dT[rsigma] = true
					break
				}
			}
			//vks := distributeVk(Vis.innerNodes[k])

			for j := 0; j < n; j++ {
				nodesInk := distributeVks[k][j]
				if len(nodesInk) > 1 && len(nodesInk) < n {
					q := len(nodesInk) - 1
					for f := 0; f < k; f++ {
						if f != sigma {
							nodesInf := distributeVks[f][j]
							for i2, n2 := range nodesInf {
								if i2 < len(nodesInf)-1 {
									Eks[k] = append(Eks[k], Edge{nodesInk[q], n2})
									dT[nodesInk[q].j] = true
									dT[n2.j] = true
								} else {
									Eks[k] = append(Eks[k], Edge{nodesInk[0], n2})
									dT[nodesInk[0].j] = true
									dT[n2.j] = true
								}
							}
						}
					}
					for f := k + 1; f < t; f++ {
						if f != sigma {
							nodesInf := distributeVks[f][j]
							for i2, n2 := range nodesInf {
								if i2 < len(nodesInf)-1 {
									Eks[k] = append(Eks[k], Edge{nodesInk[0], n2})
									dT[nodesInk[0].j] = true
									dT[n2.j] = true
								} else {
									Eks[k] = append(Eks[k], Edge{nodesInk[q], n2})
									dT[nodesInk[q].j] = true
									dT[n2.j] = true
								}
							}
						}
					}
				}

				if len(nodesInk) <= 1 {
					for g := 0; g < n; g++ {
						c := j*n + g
						if !dT[c] && !rem[c] {
							for x := 0; x <= n; x++ {
								pos := x*n + g
								if pos >= np {
									break
								}
								if rem[pos] {
									Eks[k] = append(Eks[k], Edge{node{h, i, j*n + g}, node{h, i, pos}})
									break
								}
							}
						}
					}
				}
			}
		}
	}
	return Eks
}

/*
这是将点集 按照所属不同的 Bcube(n,0)分类
*/
func distributeVk(array nodeArray) []nodeArray {
	arrays := make([]nodeArray, globeN)
	for _, n := range array {
		arrays[n.j/globeN] = append(arrays[n.j/globeN], n)
	}
	return arrays
}

/*
基本测试代码
*/
func Test(n int) {
	//准备工作
	knn := BuildCISTinKNN(n)
	buildBcube(n)

	//调用算法二生成独立生成树
	BuildCISTsInLGBCN31(n, knn)

	//检查生成树的正确性
	//CheckResult(edges)

	//打印,需要不同的类型打印主要修改这里
	// helpToFindErr(edges)

	// 打印输出
	//printToFile(n, edges)
}

func helpToFindErr(edges []Ek) {
	for k, ek := range edges {
		if k >= t {
			break
		}
		fmt.Printf("下面是第%d课独立生成树\n", k+1)
		up := make([]edgeArray, np)
		down := make([]edgeArray, np)
		for _, e := range ek.edges {
			if e.left.h == e.right.h && e.left.i == e.right.i {
				if e.left.h == 0 {
					up[e.left.i] = append(up[e.left.i], e)
				} else {
					down[e.left.i] = append(down[e.left.i], e)
				}
			}
		}
		//这里就是上面所有Bcube的生成树 按照不能 Bcube区分开了
		for i, array := range up {
			if len(array) != np-1 {
				rem := make([]bool, np)
				fmt.Printf("第%d课独立生成树,上面的 第%d个Bcube的生成树,边数量为%d\n", k+1, i, len(array))

				for _, e := range array {
					rem[e.left.j] = true
					rem[e.right.j] = true
					//fmt.Printf("%d %d\n", e.left.j, e.right.j)
				}
				for i2, _ := range rem {
					if !rem[i2] {
						fmt.Printf("%d  ", i2)
					}
				}
				fmt.Println()
			}
		}
		//这里就是下面所有Bcube的生成树 按照不能 Bcube区分开了
		for i, array := range down {
			if len(array) != np-1 {
				rem := make([]bool, np)
				fmt.Printf("第%d课独立生成树,下面的 第%d个Bcube的生成树,边数量为%d\n", k+1, i, len(array))
				for _, e := range array {
					rem[e.left.j] = true
					rem[e.right.j] = true
					//fmt.Printf("%d %d\n", e.left.j, e.right.j)
				}
				for i2, _ := range rem {
					if !rem[i2] {
						fmt.Printf("%d  ", i2)
					}
				}
				fmt.Println()
			}
		}
	}
}

func CheckResult(edges []Ek) {
	//打印,需要不同的类型打印主要修改这里
	for k, ek := range edges {
		up := make([]edgeArray, np)
		down := make([]edgeArray, np)
		for _, e := range ek.edges {
			if e.left.h == e.right.h && e.left.i == e.right.i {
				if e.left.h == 0 {
					up[e.left.i] = append(up[e.left.i], e)
				} else {
					down[e.left.i] = append(down[e.left.i], e)
				}
			}
		}
		//这里就是上面所有Bcube的生成树 按照不能 Bcube区分开了
		for i, array := range up {
			if len(array) != np-1 {
				fmt.Printf("第%d课独立生成树,上面的 第%d个Bcube的生成树,边数量为%d\n", k+1, i, len(array))
				fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			}
			rem := make([]bool, np)
			for _, e := range array {
				rem[e.left.j] = true
				rem[e.right.j] = true
			}
			for i2, _ := range rem {
				if !rem[i2] {
					fmt.Printf("%d  ", i2)
				}
			}
		}
		//这里就是下面所有Bcube的生成树 按照不能 Bcube区分开了
		for i, array := range down {
			//if len(array) != np-1 {
			if len(array) != np-1 {
				fmt.Printf("第%d课独立生成树,下面的 第%d个Bcube的生成树,边数量为%d\n", k+1, i, len(array))
				fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			}
			rem := make([]bool, np)

			for _, e := range array {
				rem[e.left.j] = true
				rem[e.right.j] = true
				//fmt.Printf("%d %d\n", e.left.j, e.right.j)
			}
			for i2, _ := range rem {
				if !rem[i2] {
					fmt.Printf("%d  ", i2)
				}
			}
		}
	}
	//用来记录有没有重复的边
	rem := make(map[string]bool)
	for i := 0; i < t; i++ {
		for _, e := range edges[i].edges {
			edg := fmt.Sprintf("%d-%d-%d %d-%d-%d", e.left.h, e.left.i, e.left.j, e.right.h, e.right.i, e.right.j)
			if rem[edg] {
				fmt.Println(edg + "边重了")
			}
		}
	}
}

func printToFile(n int, edges []Ek) {

	//创建输出文件
	fileName := "./data" + strconv.Itoa(n) + ".txt"
	file, err := os.OpenFile(fileName, syscall.O_CREAT|syscall.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Print(err)
		return
	}

	//打印所有的边
	for i := 0; i < t; i++ {
		_, _ = fmt.Fprintf(file, "begin\n")
		for _, e := range edges[i].edges {
			fmt.Fprintf(file, "%d-%d-%d %d-%d-%d\n", e.left.h, e.left.i, e.left.j, e.right.h, e.right.i, e.right.j)
		}
	}
	fmt.Fprintf(file, "end\n")
	fmt.Printf("已生成。。。\n请查看文件%s\n", fileName[2:])
	tmp := 0
	fmt.Scan(&tmp)
}

func TestSort() {
	array := nodeArray{node{0, 1, 2}, node{0, 1, 3}}
	sort.Slice(array, func(i, j int) bool {
		return array[i].j < array[j].j
	})
	fmt.Println(array)
}

func TestVPath() {
	var a1 nodeArray
	var a2 nodeArray

	for i := 0; i < 9; i++ {
		a1 = append(a1, node{0, 0, i})
	}
	a2 = append(a2, node{0, 0, 9})

	n1 := InnerNodeArraysInOneBCubeN1{false, 0, []nodeArray{a1, a2}}

	globeN = 4
	np = 16
	t = 2
	w := getW(n1)
	//fmt.Println(w)
	distributeNode(&n1, w, 0, globeN, 0, 0)
	//

	path := VPath(0, n1)
	fmt.Println(path)
	//fmt.Print(n1)

}

func TestDele() {
	var a1 nodeArray

	for i := 0; i < 9; i++ {
		a1 = append(a1, node{0, 0, i})
	}
	deleteNode(&a1, 3)
	fmt.Println(a1)
}

func TestBuildBcube() {
	buildBcube(12)
	for _, ints := range BcuMap {
		fmt.Println(ints)
	}
}
