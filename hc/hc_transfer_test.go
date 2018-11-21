/**
  * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
  *  
  * Copyright © 2015--2018 . All rights reserved.
  *
  * File: hc_transfer.go, Date: 2018-10-12
  *
  *
  * This library is free software under the terms of the GNU General Public License 
  * as published by the Free Software Foundation; either version 3 of the License, 
  * or (at your option) any later version.
  *
  */

package hc

import (
	"fmt"
	"testing"
)

func TestHC_getinfo(t *testing.T) {
	HC_getinfo()
}


func TestHC_sendtoaddress(t *testing.T) {
	HC_sendtoaddress("Tsmiz6QccgFct693jxZN1qGNVa6jaeKk6jd", 3.5)
}

func TestHC_gettransaction(t *testing.T) {
	HC_gettransaction("065a43d3fb38be18b018b1f404e981caab686aa12306167be2c1e9c5f6060722")
}

func TestHC_sendtomany(t *testing.T) {

}

func TestCreate_testaddr(t *testing.T) {
	Create_testaddr()
}

func TestRead_testaddr(t *testing.T) {

	fmt.Println(Read_testaddr("./testaddr"))

}

func TestDespatch_reward(t *testing.T) {
	//Despatch_reward("./address")
}