package Array

import "testing"

func TestSumEvenAfterQueries(t *testing.T) {
	A := []int{1,2,3,4}
	queries := [][]int{{1,0},{-3,1},{-4,0},{2,3}}
	t.Log(sumEvenAfterQueries(A, queries))
}
