package main

//Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// Leetcode 690  5.1
func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	for _, employee := range employees {
		m[employee.Id] = employee
	}
	return getImportanceByMap(m, id)
}

func getImportanceByMap(m map[int]*Employee, id int) int {
	tmp := m[id]
	var ans = tmp.Importance
	for _, subordinate := range tmp.Subordinates {
		ans += getImportanceByMap(m, subordinate)
	}
	return ans
}

// 554砖墙  2021 - 5.2
func leastBricks(wall [][]int) int {
	rem := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		tmp := 0
		for j := 0; j < len(wall[i])-1; j++ {
			tmp += wall[i][j]
			rem[tmp]++
		}
	}
	var ans = len(wall)
	var maxSize = 0
	for _, size := range rem {
		if size > maxSize {
			maxSize = size
		}
	}
	return ans - maxSize
}
