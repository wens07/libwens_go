/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: address_test
 * Date: 2018-09-04
 *
 */

package hx

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {

	seed := MnemonicToSeed("venture lazy digital aware plug hire acquire abuse chunk know gloom snow much employ glow rich exclude allow", "123")
	addr, _, _ := GetAddress(seed, "mainnet", 0, 1, 0x35)
	wif, _ := ExportWif(seed, 0, 1)

	fmt.Println(addr)
	fmt.Println(wif)

}

func TestGetAddressBytes(t *testing.T) {
}
