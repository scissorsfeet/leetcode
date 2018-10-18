package Stack

import "testing"

func TestNextGreaterElement(t *testing.T) {
	findNums := []int{4,1,2}
	nums := []int{1,3,4,2}
	t.Log(nextGreaterElement(findNums, nums))

	findNums = []int{2,4}
	nums = []int{1,2,3,4}
	t.Log(nextGreaterElement(findNums, nums))
}

