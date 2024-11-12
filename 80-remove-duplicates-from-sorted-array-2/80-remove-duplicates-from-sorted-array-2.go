package leetcode

// Problem: #80 Remove duplicates from sorted array 2
// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
// Lessons lerned:
// - if you can optimize it, optimize it (e.g. break the loop)
func removeDuplicates(nums []int) int {
	mmap := make(map[int]int)
	count := 0

	for i := 0; i < len(nums); i++ {
		if mmap[nums[i]] < 2 {
			mmap[nums[i]] = mmap[nums[i]] + 1
			nums[count] = nums[i]
			count++
		}
	}

	return count
}
