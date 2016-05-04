package MoveZeroes

func moveZeroes(nums []int) {
	numLen := len(nums)
	if numLen <= 0 {
		return
	}
	var i, j int
	var firstFind bool
	for i < numLen && j < numLen {
		if nums[i] != 0 {
			i++
		} else {
			if !firstFind {
				j = i + 1
				firstFind = true
			}
			for j < numLen && nums[j] == 0 {
				j++
			}
			if j < numLen {
				nums[i] = nums[j]
				nums[j] = 0
				i++
				j++
			}
		}
	}
}
