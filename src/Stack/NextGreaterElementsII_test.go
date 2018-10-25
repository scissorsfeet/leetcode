package Stack

import "testing"

func TestA(t *testing.T) {
	nums := []int{1,3,4,2}
	t.Log(nextGreaterElements(nums))

	nums = []int{1,2,3,4}
	t.Log(nextGreaterElements(nums))

	nums = []int{6,1,3,2,5,1}
	t.Log(nextGreaterElements(nums))
}
