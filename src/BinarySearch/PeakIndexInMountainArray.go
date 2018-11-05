package BinarySearch

func peakIndexInMountainArray(A []int) int {
	low := 0
	high := len(A)-1
	return helper(A, low, high)
}

func helper(A []int, low, high int) int {
	pivot := (low + high)/2
	if pivot == 0 || pivot == len(A)-1 {
		return 0
	}
	if A[pivot] > A[pivot-1] && A[pivot] > A[pivot+1] {
		return pivot
	}
	if high <= low {
		return 0
	}
	i := helper(A, low, pivot)
	if i > 0 {
		return i
	}
	i = helper(A, pivot+1, high)
	if i > 0 {
		return i
	}
	return 0
}
