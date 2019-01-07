package Array

import "fmt"

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}

	var res int
	for i:=0;i<n;i++ {
		var needAdd = true
		var getOne bool
		for j:=0;j<n;j++ {
			if i == j {
				continue
			}
			if M[i][j] > 0 {
				if getOne {
					helper(M, i, j, n)
				} else {
					if M[i][j] == 1 {
						helper(M, i, j, n)
						res++
						getOne = true
					}
					needAdd = false
				}
			}
		}
		if needAdd {
			res++
		}
	}

	fmt.Println(M)

	return res
}

func helper(M [][]int, i, j, n int) {
	M[i][j] = 2
	M[j][i] = 2

	for p := 0; p < n; p++ {
		if j == p {
			continue
		}
		if M[j][p] == 1 {
			helper(M, j, p, n)
		}
	}
}
