/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: serialize.go, Date: 2018-09-07
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package bts2_hdwallet

import (
	"encoding/binary"
)

// inferface for serialize hx transaction
type HxSearilze interface {
	Serialize() []byte
}

/**
 *  some basic type serialization function
 */
//func PackUint32(writer *bytes.Buffer, val uint32) ([]byte, error) {
//
//	uint64_val := uint64(val)
//
//	for {
//		uint8_val := uint8(uint64_val) & 0x7F
//
//		uint64_val >>= 7
//
//		if uint64_val > 0 {
//			uint8_val |= 0x1 << 7
//		} else {
//			uint8_val |= 0x0 << 7
//		}
//
//		err := writer.WriteByte(uint8_val)
//		if err != nil {
//			return nil, fmt.Errorf("in PackUint32 function, write byte failed: %v", err)
//		}
//
//		if uint64_val == 0 {
//			break
//		}
//
//	}
//
//	return writer.Bytes(), nil
//
//}
//
//
//func UnPackUint32(reader *bytes.Reader) (uint32, error) {
//
//	var uint32_val uint32 = 0
//	var by uint8 = 0
//	for {
//		uint8_val, err := reader.ReadByte()
//		if err != nil {
//			return 0, fmt.Errorf("in UnPackUint32 function, read byte failed: %v", err)
//		}
//
//		uint32_val |= uint32(uint8_val & 0x7F) << by
//
//		by += 7
//
//		if (uint8_val & 0x80) == 0 {
//			break
//		}
//
//	}
//
//	return uint32_val, nil
//}

func PackUint16(val uint16, isLittleEndian bool) []byte {

	res := make([]byte, 2)

	if isLittleEndian {
		binary.LittleEndian.PutUint16(res, val)
	} else {
		binary.BigEndian.PutUint16(res, val)
	}

	return res

}

func UnPackUint16(bytes []byte, isLittleEndian bool) uint16 {

	var res uint16

	if isLittleEndian {
		res = binary.LittleEndian.Uint16(bytes)
	} else {
		res = binary.BigEndian.Uint16(bytes)
	}

	return res
}

func PackUint32(val uint32, isLittleEndian bool) []byte {

	res := make([]byte, 4)

	if isLittleEndian {
		binary.LittleEndian.PutUint32(res, val)
	} else {
		binary.BigEndian.PutUint32(res, val)
	}

	return res

}

func UnPackUint32(bytes []byte, isLittleEndian bool) uint32 {

	var res uint32

	if isLittleEndian {
		res = binary.LittleEndian.Uint32(bytes)
	} else {
		res = binary.BigEndian.Uint32(bytes)
	}

	return res
}

func PackInt64(val int64, isLittleEndian bool) []byte {

	res := make([]byte, 8)

	if isLittleEndian {
		binary.LittleEndian.PutUint64(res, uint64(val))
	} else {
		binary.BigEndian.PutUint64(res, uint64(val))
	}

	return res
}

func UnPackInt64(bytes []byte, isLittleEndian bool) int64 {

	var res int64

	if isLittleEndian {
		res = int64(binary.LittleEndian.Uint64(bytes))
	} else {
		res = int64(binary.BigEndian.Uint64(bytes))
	}

	return res
}

func (asset *Asset) Serialize() []byte {

	byte_int64 := PackInt64(asset.Hx_amount, true)

	//byte for asset_id_type, default to zero
	byte_int64 = append(byte_int64, byte(0))

	return byte_int64
}

func (memo *Memo) Serialize() []byte {

	if memo.IsEmpty {
		return []byte{0}
	} else {

		//byte for optional, have element default to one
		var res []byte
		res = append(res, byte(1))
		byte_pub := make([]byte, 74)
		res = append(res, byte_pub...)
		// memo message
		res = append(res, byte(len(memo.Message)+4))
		byte_pub = make([]byte, 4)
		res = append(res, byte_pub...)
		res = append(res, []byte(memo.Message)...)
		return res

	}

}

func (tranferOp *TransferOperation) Serialize() []byte {

	res := tranferOp.Hx_fee.Serialize()
	byteTmp := make([]byte, 3)
	res = append(res, byteTmp...)

	byteTmp, _ = GetAddressBytes(tranferOp.Hx_from_addr)
	res = append(res, byteTmp...)
	byteTmp, _ = GetAddressBytes(tranferOp.Hx_to_addr)
	res = append(res, byteTmp...)

	byteTmp = tranferOp.Hx_amount.Serialize()
	res = append(res, byteTmp...)

	byteTmp = tranferOp.Hx_memo.Serialize()
	res = append(res, byteTmp...)
	res = append(res, byte(0))

	return res

}

func (trx *TransferTransaction) Serialize() []byte {

	var res []byte
	res = append(res, PackUint16(trx.Hx_ref_block_num, true)...)
	res = append(res, PackUint32(trx.Hx_ref_block_prefix, true)...)
	res = append(res, PackUint32(trx.Expiration, true)...)

	//operations
	res = append(res, byte(len(trx.Operations)))
	res = append(res, byte(0))
	for _, v := range trx.Operations {
		res = append(res, v.Serialize()...)
	}

	//extension
	res = append(res, byte(0))

	//signature
	if len(trx.Hx_signatures) > 0 {
		res = append(res, byte(len(trx.Hx_signatures)))
	}

	return res
}
