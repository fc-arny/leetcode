package leetcode

// Problem: Merge Sorted Array
// Link: https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
// Lessons lerned:
// - slices passed by ref, but when you use append it cant create a new slice
// - read the problem carefully
func merge(nums1 []int, m int, nums2 []int, n int)  {
	n_copy := append([]int(nil), nums1[:m]...)
	p1, p2 := 0, 0

	for p := 0; p < n + m; p++ {
		if p2 >= n || (p1 < m && n_copy[p1] < nums2[p2]) {
			nums1[p] = n_copy[p1]
			p1++
			continue
		}

		nums1[p] = nums2[p2]
		p2++
	}
}

func MergeProblem88(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}