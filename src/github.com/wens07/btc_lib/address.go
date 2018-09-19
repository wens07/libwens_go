/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: ecdsa.go, Date: 2018-09-19
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package btc_lib

import "github.com/btcsuite/btcd/btcec"

func GetSecp256k1Params() *btcec.KoblitzCurve {

	return btcec.S256()

}
