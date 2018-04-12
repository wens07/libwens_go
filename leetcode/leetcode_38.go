/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: leetcode_38.go
 * Date: 2018-04-12
 *
 */

package leetcode

import "github.com/wens07/algorithm"

func Leetcode_countAndSay(n int) string {

	if n == 0 {
		return ""
	}

	var res string = "1"

	for i := 1; i < n; i++ {

		var cur string = ""

		for index := 0; index < len(res); index++ {

			var count int = 1
			var j int = index

			for j+1 < len(res) && res[j] == res[j+1] {
				j++
				count++
			}

			cur += algorithm.Int2string(count) + algorithm.Int2string(int(res[j]-'0'))
			index = j

		}

		res = cur

	}

	return res

}
