/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: my_go.go
 * Date: 2018-04-11
 *
 */

package main

import (
	"fmt"
	"libwens_go/util_lib"
)

////#include <stdio.h>
////
//// void my_c_print(char* s) {
////     printf(s);
//// }
////
//import "C"



func main() {

	src := []byte{1, 2, 3, 4}

	out := util_lib.ReverseSlice(src)

	fmt.Println(out)

}
