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
	"bufio"
	"fmt"
	"os"
)

////#include <stdio.h>
////
//// void my_c_print(char* s) {
////     printf(s);
//// }
////
//import "C"

func TestVariableFunc(arg ...int) {

	fmt.Printf("%T\n", arg)

	for _, v := range arg {
		fmt.Println(v)
	}
}

func main() {

	//TestVariableFunc(1, 2, 3

	myscanner := bufio.NewScanner(os.Stdin)

	for myscanner.Scan() {
		fmt.Println(myscanner.Text())
	}
}
