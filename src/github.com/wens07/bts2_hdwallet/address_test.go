/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: address_test.go, Date: 2018-09-04
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package bts2_hdwallet

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {

	seed := MnemonicToSeed("venture lazy digital aware plug hire acquire abuse chunk know gloom snow much employ glow rich exclude allow", "123")
	addr, _ := GetAddress(seed, 0, 0)
	wif, _ := ExportWif(seed, 0, 0)

	fmt.Println(addr)
	fmt.Println(wif)

}
