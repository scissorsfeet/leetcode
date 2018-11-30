package Hash

import (
	"math"
)

func countPrimes(n int) int {
	list := make([]int, n)
	end := math.Sqrt(float64(n))
	for i:=2;float64(i)<end;i++ {
		for j := i*i; j < n; j=j+i {
			if list[j] == 0 {
				list[j] = 1
			}
		}
	}

	count := 0
	for i:=2;i<n;i++ {
		if list[i] == 0 {
			count++
		}
	}

	return count
}