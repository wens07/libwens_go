/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_832.go
 * Date: 2018-07-11
 *
 */

package leetcode

func flip(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func invert(arr []int) {
	for i, v := range arr {
		if v == 0 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}
}

func FlipAndInvertImage(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}

	for i := 0; i < len(A); i++ {
		flip(A[i])
		invert(A[i])

	}

	return A

}
