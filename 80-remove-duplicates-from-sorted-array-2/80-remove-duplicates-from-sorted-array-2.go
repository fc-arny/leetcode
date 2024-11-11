package leetcode

// Problem: #80 Remove duplicates from sorted array 2
// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicates(nums []int) int {
	i, d := 0, 1
	for i = 0; i < len(nums); i++ {
		if i <= 0 {
			continue
		}

		if nums[i] != nums[i-1] {
			d = 1
		} else {
			d++
		}

		if nums[i] == nums[i-1] && d > 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return i
}
