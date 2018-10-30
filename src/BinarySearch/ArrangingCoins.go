package BinarySearch

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}

	low := 1
	high := n
	res := 0
	for low <= high {
		mid := (low + high) >> 1
		tmp := ((1 + mid) * (mid)) >> 1
		if tmp > n {
			high = mid - 1
		} else if tmp < n {
			low = mid + 1
		} else {
			res = mid
			break
		}
	}

	if res == 0 {
		res = low - 1
	}


	return res
}