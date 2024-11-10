package leetcode

// Problem: Remove elemtn
// Link: https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
// Lessons lerned:
// - Using AI for programming - the out future (easy generating test)
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func RemoveElement27Problmen(nums []int, val int) int {
	return removeElement(nums, val)
}
