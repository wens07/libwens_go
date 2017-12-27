/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2017 . All rights reserved.
 */

package main

import (
	"fmt"

	"github.com/wens07/eth_lib"
)

func main() {

	//fmt.Println(eth_lib.WEB3_clientVersion())
	res := eth_lib.ETH_check_transaction_successful("0xe7c41a57d6722bd01c26e69b70c72119f1aacbeec18cc070cc981e986ae34adc")
	fmt.Println(res)
}
