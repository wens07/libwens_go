/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_58.go
 * Date: 2018-04-12
 *
 */

package leetcode

import (
	"fmt"
	"strings"
)

func Leetcode_lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	s = strings.TrimRight(s, " ")
	fmt.Println(s)
	var blank_index int = strings.LastIndexByte(s, ' ')

	if blank_index == len(s) {
		return 0
	}

	return len(s[blank_index+1:])

}
