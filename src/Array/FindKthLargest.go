package Array

func findKthLargest(nums []int, k int) int {
	res := helper(nums, 0, len(nums)-1, k)
	return res
}

func helper(nums []int, low, high, k int) int {
	if low > high {
		return -1
	}
	pivot := partition(nums, low, high)
	diff := (len(nums)-1-pivot) - (k-1)
	if diff == 0 {
		return nums[pivot]
	} else if diff > 0 {
		return helper(nums, pivot+1, high, k)
	} else {
		return helper(nums, low, pivot, k)
	}
}

//func partition(nums []int, low, high int) int {
//	key := nums[low]
//
//	for low < high {
//		for low < high && nums[high] > key {
//			high--
//		}
//		nums[low] = nums[high]
//
//		for low < high && nums[low] < key {
//			low ++
//		}
//		nums[high] = nums[low]
//		fmt.Println(low, high)
//	}
//
//	nums[low] = key
//
//	fmt.Println(nums)
//
//	return low
//}

func partition(arr []int, low, high int) int {
	pivotV := arr[low]
	for low < high {
		for low < high && arr[high] >= pivotV { //指针从右边开始向右找到一个比pivot小的数
			high--
		}
		arr[low] = arr[high] //将这个数放到low位，注意第一次这个位置放的是pivot值，所以不会丢

		for low < high && arr[low] <= pivotV { //指针从左边开始向右找到第一个比pivot大的数
			low++
		}
		arr[high] = arr[low] //将这个数赋值给之前的high指针，因为之前high指针指向的数已经被一定，所以不会丢
	}

	//最后将pivot的值放入合适位置，此时low与high相等
	arr[low] = pivotV
	return low
}
