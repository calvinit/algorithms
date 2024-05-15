package main

import "fmt"

// 26. 删除有序数组中的重复项: https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
//
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次，返回删除后数组的新长度。
// 元素的 相对顺序 应该保持 一致。然后返回 nums 中唯一元素的个数。
//
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
//   - 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
//   - 返回 k 。
func removeDuplicates(nums []int) int {
	length := len(nums)
	// 一个或一个元素以下的结果就是其本身
	if length <= 1 {
		return length
	}
	k := 0 // 左指针（含）的左边的元素均是不重复的
	for i := 1; i < length; i++ {
		if nums[k] != nums[i] {
			k++               // 先左指针右移一位
			nums[k] = nums[i] // 然后再把右指针位置上不存在的值填充到左指针所在的这个位置上
		}
	}
	// 注意需要返回的是长度
	return k + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums)) // expect: 2, nums = [1,2,_]
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums)) // expect: 5, nums = [0,1,2,3,4,_,_,_,_,_]
}
