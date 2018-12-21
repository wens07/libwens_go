/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: websocket.go, Date: 2018-12-19
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package rpc

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const (
	USERNAME = "test"
	PASSWD   = "test"
	HOSTURL  = "192.168.1.124"
)

func CallRpc(rpc ...string) []byte {

	u := url.URL{Scheme: "ws", Host: HOSTURL, User: url.UserPassword(USERNAME, PASSWD)}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	return nil
}
