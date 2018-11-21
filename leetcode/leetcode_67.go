/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_67.go
 * Date: 2018-04-13
 *
 */

package leetcode

import (
	"github.com/wens07/algorithm"
)

func helper_addBinary(a string, b string) string {

	res := make([]rune, len(a)+1)

	var index int = 0
	var carry uint8 = 0
	for ; index < len(b); index++ {

		tmpsum := (a[index] - '0') + (b[index] - '0') + carry
		carry = tmpsum / 2
		res[index] = rune(tmpsum%2 + '0')

	}

	for index < len(a) {

		tmpsum := (a[index] - '0') + carry
		carry = tmpsum / 2
		res[index] = rune(tmpsum%2 + '0')
		index++
	}

	if carry > 0 {
		res[index] = rune(carry + '0')
	} else {
		res = res[0:len(a)]
	}

	return algorithm.Str_reverse(string(res))
}

func Leetcode_addBinary(a string, b string) string {

	if len(a) > len(b) {

		return helper_addBinary(algorithm.Str_reverse(a), algorithm.Str_reverse(b))
	} else {
		return helper_addBinary(algorithm.Str_reverse(b), algorithm.Str_reverse(a))
	}

}
