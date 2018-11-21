/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_125.go
 * Date: 2018-04-13
 *
 */

package leetcode

import (
	"strings"

	"github.com/wens07/algorithm"
)

func get_alphanumeric_string(src string) string {

	var result string = ""
	for _, val := range src {
		if algorithm.Is_alphanumeric(string(val)) {
			result += string(val)
		}
	}

	return strings.ToLower(result)

}

func Leetcode_isPalindrome(src string) bool {

	src = get_alphanumeric_string(src)

	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {

		if src[i] != src[j] {
			return false
		}
	}

	return true

}
