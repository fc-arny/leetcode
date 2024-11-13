package leetcode

// Problem: 169. Majority Element
// Link: https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
// Lessons lerned:
// - Now I know interesting algorithm for finding majority -Boyer–Moore, see https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElement(nums []int) int {
	mmap, l := make(map[int]int), len(nums)

	major := 0
	for _, v := range nums {
		mmap[v] = mmap[v] + 1
		if mmap[v] > l/2 {
			major = v
			break
		}
	}

	return major
}
