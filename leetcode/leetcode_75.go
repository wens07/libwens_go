/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_75.go
 * Date: 2018-07-18
 *
 */

package leetcode

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func SortColors(nums []int) {

	var index_0 int = -1
	var index_1 int = -1
	var index_2 int = -1

	for i, v := range nums {

		index_1 = max(index_0, index_1)
		index_2 = max(max(index_1, index_0), index_2)

		if v == 0 {
			index_0++
			nums[index_0], nums[i] = nums[i], nums[index_0]
			if nums[i] == 1 {
				index_1++
				nums[index_1], nums[i] = nums[i], nums[index_1]
			}

		} else if v == 1 {
			index_1++
			nums[index_1], nums[i] = nums[i], nums[index_1]
		} else if v == 2 {
			index_2++
			nums[index_2], nums[i] = nums[i], nums[index_2]
		}
	}

}
