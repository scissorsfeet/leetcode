package BinarySearch

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)

	low := 0
	high := n - 1
	var index = -1
	for low <= high {
		mid := (low + high) / 2
		if letters[mid] == target {
			index = mid
			break
		} else if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	var res byte
	if index == -1 {
		if low <= n - 1 {
			for i:=low;i<n;i++ {
				if letters[i] != target {
					res = letters[i]
					break
				}
			}
		}
	} else {
		if index != n - 1 {
			for i:=index;i<n;i++ {
				if letters[i] != target {
					res = letters[i]
					break
				}
			}
		}
	}
	if res == 0 {
		res = letters[0]
	}
	return res
}
