/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: trx_test.go
 * Date: 2018-09-05
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

	transferTrx := Transaction{
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

	//addrKey, _ := GetAddressKey(seed, 0, 0)

	wif, _ := ExportWif(seed, 0, 0)

	sig, _ := GetSignature(wif, []byte(chainidHex))

	fmt.Println(len(sig))
	fmt.Println(sig)

}

func TestTransaction(t *testing.T) {

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
	//out, err := BuildTransferTransaction("20295,1522980328", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq", "HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", "HXNUFt3gViQn3iXLRL9eTMGsaJFhYQ6RT4gM",
	//	"", 10000000, 200, "HC", "", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")
	//ref_block_num, ref_block_prefix, _ := GetRefblockInfo("47926,2298351426")
	//
	//fmt.Printf("ref block num is %d\n", ref_block_num)
	//fmt.Printf("ref block prefix is %d\n", ref_block_prefix)

	//out, err := BuildBindAccountTransaction("6413,2521010061", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq", "HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", 0, "TshDfDSPRhV2BCDAJFAGAjs5K2TJGfrCPua",
	//	"HC", "PtWVJdsvDGYsidC9igf6h2KRBFCdUb7k6Phx9DoZPeXzHirhL2yAM", "", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")

	//BuildUnBindAccountTransaction("6413,2521010061", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq","HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", 0, "TshDfDSPRhV2BCDAJFAGAjs5K2TJGfrCPua",
	//	"HC", "PtWVJdsvDGYsidC9igf6h2KRBFCdUb7k6Phx9DoZPeXzHirhL2yAM", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")

	//BuildWithdrawCrosschainTransaction("39618,358453409", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq","HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", 0, "TshDfDSPRhV2BCDAJFAGAjs5K2TJGfrCPua",
	//	"HC", "1.2", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")

	//BuildRegisterAccountTransaction("39618,358453409", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq","HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", "HX77DEz5FFdsbyM4P4XMyZ5Xm2DHPph4o3GjLXcyc8Eq62s84SMw",500000, "", "wens", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")

	//out, err := BuildLockBalanceTransaction("33081,1272682105", "5KR6ocp5eUdWWYPX7mYp4XLGBcZ2xHVHVsNaco6K2YZSWQTqES7", "HXNcikaxB2rsK26JCiwvzse9AqFPkBGyAynG", "1.2.60", "1.3.0", 100000, 0,
	//	"1.6.11", "InvalidAddress", "07c870b857439cc298de0f7747d475c57320ddfdd6f28357f7bed2a7ff41e821")

	out, err := BuildRedeemBalanceTransaction("59990,2111459066", "5KR6ocp5eUdWWYPX7mYp4XLGBcZ2xHVHVsNaco6K2YZSWQTqES7", "HXNcikaxB2rsK26JCiwvzse9AqFPkBGyAynG", "1.2.60", "1.3.0", 200000, 0,
		"1.6.13", "InvalidAddress", "07c870b857439cc298de0f7747d475c57320ddfdd6f28357f7bed2a7ff41e821")

	//asset_arr := []string{"citizen10,54459861,1.3.0", "citizen9,39886,1.3.0"}
	//out, err := BuildObtainPaybackTransaction("5595,4227186882", "5KbxwmcNhUQe7oVN5oMpC3BiYmGpNDf8u3W1EPn3qzfrGxVahyq", "HXNRhTbDKiw2ut91BJEc5zy49HKWQnRw9as7", 200, asset_arr,
	//	"2.22.15", "9f3b24c962226c1cb775144e73ba7bb177f9ed0b72fac69cd38764093ab530bd")
	//

	//out, err := BuildContractInvokeTransaction("44746,1493958321", "5KR6ocp5eUdWWYPX7mYp4XLGBcZ2xHVHVsNaco6K2YZSWQTqES7", "HXNcikaxB2rsK26JCiwvzse9AqFPkBGyAynG", 10100,
	//	100, 10000, "HXCKkb72yRQ16fkhA9pGvX2p95qAhsADZTK2", "transfer", "HXNLi2d4m6gsRdsL3FQbBBeCLJEU6QRotg8n,15", "", "07c870b857439cc298de0f7747d475c57320ddfdd6f28357f7bed2a7ff41e821")

	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(string(out))

}
