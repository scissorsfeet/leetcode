package MinimumMovesToEqualArrayElements

/*
Given a non-empty integer array of size n,
find the minimum number of moves required to make all array elements equal,
where a move is incrementing n - 1 elements by 1.
*/

func minMoves(nums []int) int {
    n := len(nums)
    if n == 0{
        return 0
    }
    min := nums[0]
    var sum int
    for i,_ := range nums{
        if nums[i] < min {
            min = nums[i]
        }
        sum += nums[i]
    }

    return sum-n*min
}
