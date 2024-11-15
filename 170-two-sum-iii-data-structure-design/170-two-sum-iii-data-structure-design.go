package leetcode

import (
	"sort"
)

// Problem: 170. Two Sum III - Data structure design
// Link: https://leetcode.com/problems/two-sum-iii-data-structure-design/
// Lessons lerned:
// - I remembered how to create structures
type TwoSum struct {
	arr []int
}

func Constructor() TwoSum {
	return TwoSum{[]int{}}
}

func (this *TwoSum) Add(number int) {
	this.arr = append(this.arr, number)
}

func (this *TwoSum) Find(value int) bool {
	i, j := 0, len(this.arr)-1

	sort.Ints(this.arr)
	for i < j {
		if this.arr[i]+this.arr[j] < value {
			i++
		} else if this.arr[i]+this.arr[j] > value {
			j--
		} else {
			return true
		}
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
