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
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
)

const api_url string = "https://127.0.0.1:12010/rpc"


func Rpc_call(args ...string) *simplejson.Json {

	id := "1"
	method := args[0]
	params := "[]"
	if len(args) > 1 {
		params = args[1]
	}

	postBody := `{"jsonrpc":"2.0","method":"` + method + `","params":` + params + `,"id":` + id + `}`
	fmt.Println("postBody: " + postBody)

	_, body, errs := gorequest.New().TLSClientConfig(&tls.Config{ InsecureSkipVerify: true}).Post(api_url).SetBasicAuth("test", "test").
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

func HC_getinfo()  {
	res := Rpc_call("getinfo").Get("result").MustMap()
	fmt.Println(reflect.TypeOf(res))
	fmt.Println(res)
}

func HC_sendtoaddress(addr string, amount float64)  {

	res := Rpc_call("sendtoaddress", `["` + addr + `", ` + strconv.FormatFloat(amount, 'f', -1, 64) + `]`).Get("result").MustString()
	fmt.Println(res)
	
}

func HC_sendtomany(account string, param string) {

	res := Rpc_call("sendtomany", `["` + account + `", ` + param + `]`).Get("result").MustString()
	fmt.Println(res)
}


func HC_gettransaction(addr string) {
	res := Rpc_call("gettransaction", `["` + addr + `"]`).Get("result").MustMap()
	fmt.Println(res)
}




func Transfer2addr(addr string, amount string) string {
	res := Rpc_call("sendtoaddress", `["` + addr + `", ` +  amount + `]`).Get("result").MustString()
	fmt.Println(res)

	return res
}


func Create_testaddr() {

	f, err := os.OpenFile("testaddr", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}


	f.WriteString("{\n")
	addr := "Tsmiz6QccgFct693jxZN1qGNVa6jaeKk6jd"
	base := 0.01

	for i := 0; i < 400; i++ {

		amount := base * float64(i)

		f.WriteString(`"` + addr + `"` + ":"+strconv.FormatFloat(amount, 'f', 2, 64))
		f.WriteString("\n")

	}


	f.WriteString("}")


	defer f.Close()


}


func Read_testaddr(fname string) string {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0755)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return string(data)



}

func Despatch_reward(addrFile string, outTrx string) int {

	addrfile, err := os.Open(addrFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	success := 0

	defer addrfile.Close()

	scanner := bufio.NewScanner(addrfile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		line :=	scanner.Text()
		if len(line) <= 35 {
			panic("the line in address file error!")
		}

		transferInfo := strings.Split(line, ":")

		if len(transferInfo) != 2 {
			panic("parse address file error!")
		}

		Transfer2addr(transferInfo[0], transferInfo[1])


		success++
	}

	return success
}







