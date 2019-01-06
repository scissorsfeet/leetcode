package Array

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	t.Log(increasingTriplet([]int{1,2,3,4,5}))
	t.Log(increasingTriplet([]int{2,4,1,3,5}))
	t.Log(increasingTriplet([]int{5,4,3,2,1}))
	t.Log(increasingTriplet([]int{5,1,5,5,2,5,4}))
	t.Log(increasingTriplet([]int{0,-1,-1,1,-1,-1,2}))
}
