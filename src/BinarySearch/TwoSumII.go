package BinarySearch

func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		t := target-numbers[i]
		index := search(numbers[i+1:], t)
		if index != -1 {
			return []int{i+1, i+1+index+1}
		}
	}
	return []int{-1,-1}
}