/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: myip.go, Date: 2018-11-30
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package util_lib

import (
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func GetMyIp(apiUrl string) string {

	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("response code error!")
	}

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	return string(bodyByte)
}

func GetIpInfo(ip string) string {

	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("response code error!")
	}

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	resJson, _ := simplejson.NewJson(bodyByte)
	country := resJson.Get("data").Get("country").MustString()
	region := resJson.Get("data").Get("region").MustString()
	city := resJson.Get("data").Get("city").MustString()

	isp := resJson.Get("data").Get("isp").MustString()

	return country + "/" + region + "/" + city + " : " + isp

}
