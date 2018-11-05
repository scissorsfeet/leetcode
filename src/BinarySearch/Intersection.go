package BinarySearch

var searchChan = make(chan int, 1024)

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	quickSort(nums1, 0, len(nums1)-1)
	quickSort(nums2, 0, len(nums2)-1)

	var haystack []int
	var needle []int
	if len(nums1) > len(nums2) {
		haystack = nums1
		needle = nums2
	} else {
		haystack = nums2
		needle = nums1
	}

	go func() {
		for _, v := range needle {
			isExist := binarySearch(v, haystack, 0, len(haystack)-1)
			if isExist {
				searchChan <- v
			}
		}
		close(searchChan)
	}()

	var result = make([]int, 0, 1024)
	var rm = make(map[int]rune, 1024)
	for v := range searchChan {
		if _, ok := rm[v]; ok {
			continue
		}
		rm[v] = 1
		result = append(result, v)
	}

	return result
}

func binarySearch(needle int, haystack []int, low, high int) (bool) {
	if low < high {
		pivot := (low + high) / 2
		if haystack[pivot] == needle {
			return true
		} else if haystack[pivot] > needle {
			return binarySearch(needle, haystack, low, pivot-1)
		} else {
			return binarySearch(needle, haystack, pivot+1, high)
		}
	}
	return false
}

//快排
func quickSort(nums []int, i, j int) {
	if j <= i {
		return
	}
	low := i
	high := j
	key := nums[low]

	for low < high {
		for low < high && nums[high] >= key {
			high--

		}
		nums[low] = nums[high]
		for low < high && nums[low] <= key {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = key
	quickSort(nums, i, low-1)
	quickSort(nums, low+1, j)
}