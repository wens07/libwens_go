/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2017 . All rights reserved.
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package eth_lib

import (
	"database/sql"
	"fmt"
)

const (
	mysql_conn_str string = "wq:123456@tcp(192.168.1.123:3306)/ethAddrBalance?charset=utf8"
)

func Insert_address_byscanningblock(blocknum uint64, db *sql.DB) {

	addr_strarr := ETH_getTrxHashsByBlockNumber(blocknum)
	fmt.Println("Trx hash num: ", len(addr_strarr))
	if len(addr_strarr) == 1 && addr_strarr[0] == "" {
		return
	}

	for _, v := range addr_strarr {

		fmt.Println(v)

		addr_from, addr_to := ETH_getTransactionByHash(v)

		if addr_from != "" {
			insert_str := "insert ignore into addr_balance (addr) VALUES (" + `"` + addr_from + `"` + ")"
			//fmt.Println(insert_str)
			_, err := db.Exec(insert_str)
			CheckErr(err)
		}

		if addr_to != "" {
			insert_str := "insert ignore into addr_balance (addr) VALUES (" + `"` + addr_to + `"` + ")"
			//fmt.Println(insert_str)
			_, err := db.Exec(insert_str)
			CheckErr(err)
		}

	}

}

func Scan_block(blocknum_s uint64, blocknum_e uint64) {

	db := Connect_db(mysql_conn_str)
	defer db.Close()

	for blocknum_s <= blocknum_e {

		Insert_address_byscanningblock(blocknum_s, db)

		blocknum_s += 1

	}

}
