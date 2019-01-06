package Array

func increasingTriplet(nums []int) bool {
	numsLen := len(nums)
	if numsLen < 3 {
		return false
	}

	first := 0
	second := -1

	for i:=1;i<numsLen;i++ {
		cur := nums[i]
		if second != -1 {
			if cur > nums[second] {
				return true
			} else if cur <= nums[second] && cur > nums[first] {
				second = i
			} else if cur < nums[first] {
				first = i
			}
		} else {
			if cur > nums[first] {
				second = i
			} else {
				first = i
			}
		}
	}

	return false
}