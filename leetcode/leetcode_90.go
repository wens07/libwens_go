/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_90.go
 * Date: 2018-07-20
 *
 */

package leetcode

import (
	"sort"

	"github.com/wens07/algorithm"
)

func is_array_same(src []int, dst []int) bool {

	if len(src) != len(dst) {
		return false
	}

	for i, v := range src {
		if v != dst[i] {
			return false
		}
	}

	return true
}

func is_contain(src []int, src_set [][]int) bool {

	for _, v := range src_set {

		if is_array_same(src, v) {
			return true
		}

	}

	return false
}

func SubsetsWithDup(nums []int) [][]int {

	sort.Sort(sort.IntSlice(nums))
	var tmp = algorithm.Subset(nums)

	var res [][]int

	for _, v := range tmp {
		if !is_contain(v, res) {
			res = append(res, v)
		}
	}

	return res

}
