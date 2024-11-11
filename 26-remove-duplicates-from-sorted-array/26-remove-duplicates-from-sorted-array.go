package leetcode

// Problem: Remove duplicates from sorted arrays
// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	i := 0
	for ; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return i
}
