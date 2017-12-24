package main

import (
	"fmt"
	"reflect"

	"github.com/wens07/eth_lib"
)

func check_trx_whether_success(num uint64) {

	mysql_conn_str_tmp := "wq:123456@tcp(192.168.1.123:3306)/ethAddrBalance?charset=utf8"
	db := eth_lib.Connect_db(mysql_conn_str_tmp)
	defer db.Close()

	select_str := "select trx_id from air_drop where `check` is null limit " + fmt.Sprintf("%d", num)

	fmt.Println(select_str)

	rows, err := db.Query(select_str)
	eth_lib.CheckErr(err)

	var trx_id string
	var index int = 0

	for rows.Next() {
		index++

		if err := rows.Scan(&trx_id); err != nil {
			eth_lib.CheckErr(err)
		}

		fmt.Println(trx_id)

		eth_lib.ETH_getTransactionByHash(trx_id)

		//check := "1"
		//update_str := "update addr_balance set `check` = " + `"` + check + `"` + " where addr = " + `"` + addr + `"`
		//fmt.Println(update_str)
		//
		//_, err := db.Exec(update_str)
		//eth_lib.CheckErr(err)

	}

}

func main() {

	trxhash := "0xc35a313fd7c39d9615c1ff66ace3e6d952f3b000107520cc15866ee13248e2c9"

	callArgs := `["` + trxhash + `"]`

	res := eth_lib.Remote_rpc_call("eth_getTransactionByHash", callArgs).Get("result").Get("blockNumber").MustInt64()

	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))

	if res == 0 {
		fmt.Println("block number is nil!")
	}

}
