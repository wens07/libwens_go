/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_442.go
 * Date: 2018-07-11
 *
 */

package leetcode

func FindDuplicates(nums []int) []int {

	var res []int
	nums = append(nums, 0)
	arr_len := len(nums)

	for i := 0; i < arr_len; i++ {
		index := nums[i] % arr_len
		nums[index] = nums[index] + arr_len
	}

	for i := 0; i < arr_len; i++ {
		if nums[i]/arr_len > 1 {
			res = append(res, i)
		}
	}

	return res

}
