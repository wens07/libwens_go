/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: address_test
 * Date: 2018-09-04
 *
 */

package bts

import (
	"fmt"
	"testing"
)

func TestKey(t *testing.T) {

	//seed := MnemonicToSeed("venture lazy digital aware plug hire acquire abuse chunk know gloom snow much employ glow rich exclude allow", "123")
	//wif, _ := ExportWif(seed,  0, 0)
	//fmt.Println(wif)
	privkey, _ := ImportWif("5JQcosxpBsLGhxgdYmNM83xVF1UyBpUnhiCGMJGXPY2VzQcSKEc")

	//privkey, _ := ImportWif(wif)
	pubKeystr, _ := GetPubKeyStr(privkey)

	fmt.Println(pubKeystr)

}
