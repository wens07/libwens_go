package eth_lib

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/parnurzeal/gorequest"
)

//const api_url string = "http://192.168.1.124:8588"

const api_url string = "http://192.168.1.178:8546"

func rpc_call(args ...string) *simplejson.Json {

	id := "1"
	method := args[0]
	params := "[]"
	if len(args) > 1 {
		params = args[1]
	}

	postBody := `{"jsonrpc":"2.0","method":"` + method + `","params":` + params + `,"id":` + id + `}`
	fmt.Println("postBody: " + postBody)

	_, body, errs := gorequest.New().Post(api_url).
		Send(postBody).
		End()

	if errs != nil {
		fmt.Println(errs)
		panic(errs)
	}

	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		log.Fatalln(err)
	}

	return js
}

func Int2hexstr(i uint64) string {

	return hexutil.EncodeUint64(i)
}

func Hexstr2int(str string) uint64 {
	return hexutil.MustDecodeUint64(str)
}

func ETH_blockNumber() string {
	res := rpc_call("eth_blockNumber").Get("result").MustString()
	return res
}

func WEB3_clientVersion() string {
	res := rpc_call("web3_clientVersion").Get("result").MustString()
	return res
}

func NET_peerCount() string {
	res := rpc_call("net_peerCount").Get("result").MustString()
	return res

}

func ETH_getTrxHashsByBlockNumber_debug(num uint64) {

	hexstr := Int2hexstr(num)

	callArgs := `["` + hexstr + `", false ]`

	//res := rpc_call("eth_getBlockByNumber", callArgs).Get("result").MustMap()["transactions"]
	res := fmt.Sprintf("%s", rpc_call("eth_getBlockByNumber", callArgs).Get("result").MustMap()["transactions"])

	str_arr := strings.Split(res[1:len(res)-1], " ")
	for _, v := range str_arr {
		fmt.Println(v)
	}
	fmt.Println(reflect.TypeOf(res))

}

func ETH_getTrxHashsByBlockNumber(num uint64) []string {

	fmt.Println("current block num: ", num)

	hexstr := Int2hexstr(num)

	callArgs := `["` + hexstr + `", false ]`

	res := fmt.Sprintf("%s", rpc_call("eth_getBlockByNumber", callArgs).Get("result").MustMap()["transactions"])

	str_arr := strings.Split(res[1:len(res)-1], " ")

	return str_arr
}

func ETH_getTransactionByHash(trxhash string) (string, string) {

	callArgs := `["` + trxhash + `"]`

	res := rpc_call("eth_getTransactionByHash", callArgs).Get("result")

	res_from := res.Get("from").MustString()
	res_to := res.Get("to").MustString()

	//fmt.Println(res_from)
	//fmt.Println(res_to)

	return res_from, res_to
}

func ETH_getBalance(addr string) string {

	callArgs := `["` + addr + `", "latest" ]`

	res := rpc_call("eth_getBalance", callArgs).Get("result").MustString()

	balance, err := hexutil.DecodeBig(res)

	CheckErr(err)

	res = fmt.Sprintf("%s", balance)

	//fmt.Println(res)

	return res
}

func ETH_getBalance_by_block(addr string, blocknum uint64) string {

	hexblocknum := Int2hexstr(blocknum)
	fmt.Println(hexblocknum)

	callArgs := `["` + addr + `",` + `"` + hexblocknum + `"` + "]"

	res := Remote_rpc_call("eth_getBalance", callArgs).Get("result").MustString()

	//fmt.Println(res)

	balance, err := hexutil.DecodeBig(res)

	CheckErr(err)

	res = fmt.Sprintf("%s", balance)

	//fmt.Println(res)

	return res

}

func ETH_check_transaction_successful(trxhash string) bool {

	callArgs := `["` + trxhash + `"]`

	res := rpc_call("eth_getTransactionByHash", callArgs).Get("result").Get("blockNumber").MustInt64()

	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))

	if res == 0 {
		return false
	}

	return true

}

func ETH_estimateGas(sender_addr string, recieve_addr string, data string) uint64 {

	callArgs := `[{` + `"from": ` + `"` + sender_addr + `", ` + `"to": ` + `"` + recieve_addr + `", `
	callArgs += `"data": ` + `"` + data + `"` + "}]"

	res := rpc_call("eth_estimateGas", callArgs).Get("result").MustString()

	return Hexstr2int(res)

}

func ETH_pendingTransactions() int {

	res := rpc_call("eth_pendingTransactions").Get("result").MustArray()

	//fmt.Println(reflect.TypeOf(res))
	//fmt.Println(len(res))
	//fmt.Println(len(res))

	return len(res)

}
