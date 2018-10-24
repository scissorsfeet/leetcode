package BinarySearch

func isPerfectSquare(num int) bool {
	low := 1
	high := num

	for low <= high {
		mid := (low + high) / 2
		sqrt := mid * mid
		if num == sqrt {
			return true
		} else if sqrt > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
