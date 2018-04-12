/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: permutation.go
 * Date: 2018-04-12
 *
 */

package algorithm

import "fmt"

func print_result(src []string) {

	for i := 0; i < len(src); i++ {
		fmt.Print(src[i])
	}

	fmt.Println()
}

func Permutation(src []string, size int) {

	if size == 1 {
		print_result(src)
		return
	}

	for i := 0; i < size; i++ {
		Permutation(src, size-1)

		if size%2 == 0 {
			src[i], src[size-1] = src[size-1], src[i]
		} else {
			src[0], src[size-1] = src[size-1], src[0]
		}

	}

}
