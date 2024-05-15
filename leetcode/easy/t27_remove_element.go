package main

import "fmt"

// 27. 移除元素: https://leetcode.cn/problems/remove-element/description/
//
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}
	return l
}

func main() {
	nums, val := []int{3, 2, 2, 3}, 3
	fmt.Println(removeElement(nums, val)) // expect: 2, nums = [2,2,_,_]
	fmt.Println(nums)
	nums, val = []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	fmt.Println(removeElement(nums, val)) // expect: 5, nums = [0,1,3,0,4,_,_,_]
	fmt.Println(nums)
}
