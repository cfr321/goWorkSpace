package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	MAX := int(1e9 + 7)
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	} else {
		dp[1] = 1
	}

	for i := 1; i < len(s); i++ {
		be := s[i-1]
		if s[i] == '*' {
			dp[i+1] = (9 * dp[i]) % MAX
			if be == '1' {
				dp[i+1] = (dp[i+1] + dp[i-1]*9) % MAX
			}
			if be == '2' {
				dp[i+1] = (dp[i+1] + dp[i-1]*6) % MAX
			}
			if be == '*' {
				dp[i+1] = (dp[i+1] + dp[i-1]*15) % MAX
			}
		} else if s[i] == '0' {
			if be == '*' {
				dp[i+1] = 2 * dp[i-1] % MAX
			} else if be == '1' || be == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else {
			dp[i+1] = dp[i]
			if s[i] <= '6' {
				if be == '*' {
					dp[i+1] = (dp[i+1] + 2*dp[i-1]) % MAX
				}
				if be == '1' || be == '2' {
					dp[i+1] = (dp[i+1] + dp[i-1]) % MAX
				}
			} else {
				if be == '*' || be == '1' {
					dp[i+1] = (dp[i+1] + dp[i-1]) % MAX
				}
			}
		}
	}
	return dp[len(s)]
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	pre := make(map[int]int)
	pre[0] = 1
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}
		cur += node.Val
		res := pre[cur-targetSum]
		pre[cur]++
		res += dfs(node.Left, cur)
		res += dfs(node.Right, cur)
		pre[cur]--
		return res
	}
	return dfs(root, 0)
}
