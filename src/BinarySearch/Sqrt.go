package BinarySearch

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	low := 0
	high := x
	for low <= high {
		mid := (low + high) >> 1
		tmp := mid * mid
		if tmp > x {
			high = mid - 1
		} else if tmp < x {
			low = mid + 1
		} else {
			return mid
		}
	}

	return low - 1
}
