/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: trx_test.go, Date: 2018-09-05
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package hx

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTimeConvert(t *testing.T) {

	tmp_time := time.Now().Unix()

	fmt.Println(Time2Str(tmp_time))

}

func TestJson(t *testing.T) {

	transferOp := DefaultTransferOperation()

	transferTrx := TransferTransaction{
		1,
		2,
		"2018-09-04T08:16:25",
		[][]interface{}{{0, transferOp}},
		make([]interface{}, 0),
		[]string{"2018-09-04T08:16:25"},
		3,
		nil,
	}

	b, err := json.Marshal(transferTrx)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

func TestSignature(t *testing.T) {

	chainidHex := "fe70279c1d9850d4ddb6ca1f00c577bc2e86bf33d54fafd4c606a6937b89ae32"

	seed := MnemonicToSeed("venture lazy digital aware plug hire acquire abuse chunk know gloom snow much employ glow rich exclude allow", "123")

	addrKey, _ := GetAddressKey(seed, 0, 0)

	sig, _ := GetSignature(addrKey, []byte(chainidHex))

	fmt.Println(len(sig))
	fmt.Println(sig)

}

func TestPack(t *testing.T) {

	//tmp := HxSearilze{}
	//
	//
	//var val uint16 = 61162
	//res := tmp.PackUint16(val, true)
	//
	//fmt.Println(hex.EncodeToString(res))
	//
	//val = tmp.UnPackUint16(res, true)
	//
	//fmt.Println(val)

	//var val uint32 = 3949464256
	//
	//res := tmp.PackUint32(val, true)
	//
	//fmt.Println(hex.EncodeToString(res))

	//res, _:= hex.DecodeString("a612965b")
	//val := tmp.UnPackUint32(res, true)
	//
	//out := Time2Str(int64(val))
	//
	//fmt.Println(out)

	//int64
	//var val int64 = 1000000
	//res := tmp.PackInt64(val, true)
	//fmt.Println(hex.EncodeToString(res))
	//
	//val = tmp.UnPackInt64(res, true)
	//fmt.Println(val)

	//addrss
	//fromAddrBytes, _ := GetAddressBytes("HXNNGcPe7b3P39k5uss6Du6MqRDo34itTSi")
	//fmt.Println(hex.EncodeToString(fromAddrBytes))
	//
	//toAddrBytes, _ := GetAddressBytes("HX3uMGhuxfVLeQBxxB5PLNndV26xWqo8QiQ")
	//fmt.Println(hex.EncodeToString(toAddrBytes))

	////memo
	//memo := DefaultMemo()
	//memo.IsEmpty = false
	//memo.Message = "test to wens 10"
	//byte_memo := memo.Serialize()
	//fmt.Println(hex.EncodeToString(byte_memo))
	//
	////asset
	//asset := DefaultAsset()
	//asset.Hx_amount = 1000000
	//byte_asset := asset.Serialize()
	//fmt.Println(hex.EncodeToString(byte_asset))
	//
	////transferOP
	//transferOp := DefaultTransferOperation()
	//transferOp.Hx_amount.Hx_amount = 1000000
	//transferOp.Hx_fee.Hx_amount = 2018
	//transferOp.Hx_memo.IsEmpty = false
	//transferOp.Hx_memo.Message = "test to wens 10"
	//
	//transferOp.Hx_from_addr = "HXNNGcPe7b3P39k5uss6Du6MqRDo34itTSi"
	//transferOp.Hx_to_addr = "HX3uMGhuxfVLeQBxxB5PLNndV26xWqo8QiQ"
	//
	//byte_transferOp := transferOp.Serialize()
	//fmt.Println(hex.EncodeToString(byte_transferOp))

	//transferTrx
	BuildTransferTransaction("HXNNGcPe7b3P39k5uss6Du6MqRDo34itTSi", "HX3uMGhuxfVLeQBxxB5PLNndV26xWqo8QiQ", "test to wens 10", 100000)

}
