/*
Given a sorted array of distinct integers and a target value,
return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] < target {
		i++
	}
	return i
}