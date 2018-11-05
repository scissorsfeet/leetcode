package Sort

func sortColors(nums []int)  {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	times := 0
	for i:=0;i<numsLen; {
		if (times == 0 && nums[i] == 0) || (times == 1 && nums[i] == 1) {
			i++
			continue
		}
		for j:=numsLen-1;j>i;j-- {
			if (times == 0 && nums[i] != 0 && nums[j] == 0 ) || (times == 1 && nums[i] != 1 && nums[j] == 1) {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
				i++
				for i < numsLen {
					if (times == 0 && nums[i] == 0) || (times == 1 && nums[i] == 1) {
						i++
					} else {
						break
					}
				}
			}
		}
		times++
		if times == 2 {
			break
		}
	}
}
