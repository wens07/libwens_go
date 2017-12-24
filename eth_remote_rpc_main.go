package main

import (
	"fmt"

	"github.com/wens07/eth_lib"
)

func main() {

	//var num uint64 = 3910000
	//
	//hexstr := eth_lib.Int2hexstr(num)
	//
	//callArgs := `["` + hexstr + `", false ]`
	//
	//res := fmt.Sprintf("%s", remote_rpc_call("eth_getBlockByNumber", callArgs).Get("result").MustMap()["transactions"])
	//
	//fmt.Println(res)

	//res := remote_rpc_call("web3_clientVersion").Get("result").MustString()

	//addr_strarr := eth_lib.ETH_getTrxHashsByBlockNumber(4000000)
	//
	//fmt.Println("Trx hash num: ", len(addr_strarr))
	//if len(addr_strarr) == 1 && addr_strarr[0] == "" {
	//	return
	//}
	//
	//for _, v := range addr_strarr {
	//
	//	fmt.Println(v)
	//
	//	addr_from, addr_to := eth_lib.ETH_getTransactionByHash(v)
	//
	//	fmt.Println("from: ", addr_from)
	//	fmt.Println("to :", addr_to)
	//
	//}

	//fmt.Println(retmote_eth_getBalance("0x1e470f8a0f46e62e64a7911e7ec51000acca4b23", 62))
	//if eth_lib.ETH_pendingTransactions()-20 >= 26 {
	//	fmt.Println("cuur")
	//}

	//fmt.Println(eth_lib.ETH_pendingTransactions())
	fmt.Println(eth_lib.ETH_check_transaction_successful("0x1a75132659749910728157f4757dc06cba88f4837e9f98783b8c2b5fef688ca2"))

}
