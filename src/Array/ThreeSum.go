package Array

func threeSum(nums []int) [][]int {
	m := make(map[int][][2]int, 1024)
	numsLen := len(nums)
	for i:=0;i<numsLen-1;i++ {
		for j:=i+1;j<numsLen;j++ {
			sum := nums[i] + nums[j]
			if _, ok := m[sum]; !ok {
				m[sum] = make([][2]int, 0, 32)
			}
			m[sum] = append(m[sum], [2]int{i, j})
		}
	}

	res := make([][]int, 0, 32)
	for i:=0;i<numsLen;i++ {
		target := 0 - nums[i]
		if _, ok := m[target]; ok {
			for j, _ := range m[target] {
				if i != m[target][j][0] && i != m[target][j][1] {
					res = append(res, []int{nums[i], nums[m[target][j][0]], nums[m[target][j][1]]})
				}
			}
		}
	}

	return res
}
