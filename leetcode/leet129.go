//
// Author: cfr
//

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = 0

func sumNumbers(root *TreeNode) int {
	res = 0
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, i int) {
	if root.Right == nil && root.Left == nil {
		res += i*10 + root.Val

	}
	if root.Left != nil {
		dfs(root.Left, i*10+root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, i*10+root.Val)
	}
}
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i != 0 && grid[i-1][j] == 1 {
					res--
				}
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					res--
				}
				if j != 0 && grid[i][j-1] == 1 {
					res--
				}
				if j != len(grid[0])-1 && grid[i][j+1] == 1 {
					res--
				}
			}
		}
	}
	return res
}
