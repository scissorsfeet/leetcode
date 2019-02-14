package Array

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var evenSum int
	for _, v := range A {
		if isEven(v) {
			evenSum += v
		}
	}

	var res = make([]int, 0, len(queries))
	for _, query := range queries {
		index := query[1]
		val := query[0]
		old := A[index]
		A[index] += val
		if isEven(A[index]) {
			if isEven(old) {
				evenSum += A[index] - old
			} else {
				evenSum += A[index]
			}
		} else {
			if isEven(old) {
				evenSum -= old
			}
		}
		res = append(res, evenSum)
	}

	return res
}

func isEven(x int) bool {
	if x < 0 {
		x = 0-x
	}
	if x%2==0 {
		return true
	}
	return false
}
