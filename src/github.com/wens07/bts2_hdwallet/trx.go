/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: trx.go, Date: 2018-09-04
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package hx

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/btcsuite/btcutil/hdkeychain"
	bts "github.com/wens07/bts2_secp256k1"
)

const (
	chain_id string = "fe70279c1d9850d4ddb6ca1f00c577bc2e86bf33d54fafd4c606a6937b89ae32"
)

// define trx structure
type TransferTransaction struct {
	Hx_ref_block_num    uint16 `json:"ref_block_num"`
	Hx_ref_block_prefix uint32 `json:"ref_block_prefix"`
	Hx_expiration       string `json:"expiration"`

	Hx_operations [][]interface{} `json:"operations"`
	Hx_extensions []interface{}   `json:"extensions"`
	Hx_signatures []string        `json:"signatures"`

	Expiration uint32              `json:"-"`
	Operations []TransferOperation `json:"-"`
}

func DefaultTransferTransaction() *TransferTransaction {

	return &TransferTransaction{
		0,
		0,
		"",
		nil,
		nil,
		nil,
		0,
		nil,
	}
}

type Asset struct {
	Hx_amount   int64  `json:"amount"`
	Hx_asset_id string `json:"asset_id"`
}

func DefaultAsset() Asset {
	return Asset{
		0,
		"1.3.0",
	}
}

type Extension struct {
	extension []string
}

type Memo struct {
	Hx_from    string `json:"from"` //public_key_type  33
	Hx_to      string `json:"to"`   //public_key_type  33
	Hx_nonce   uint64 `json:"nonce"`
	Hx_message string `json:"message"`

	IsEmpty bool   `json:"-"`
	Message string `json:"-"`
}

func DefaultMemo() Memo {

	return Memo{
		"HX1111111111111111111111111111111114T1Anm",
		"HX1111111111111111111111111111111114T1Anm",
		0,
		"",
		true,
		"",
	}

}

// transfer operation tag is  0
type TransferOperation struct {
	Hx_fee  Asset  `json:"fee"`
	Hx_from string `json:"from"`
	Hx_to   string `json:"to"`

	Hx_from_addr string `json:"from_addr"`
	Hx_to_addr   string `json:"to_addr"`

	Hx_amount Asset `json:"amount"`
	Hx_memo   Memo  `json:"memo"`

	Hx_extensions []interface{} `json:"extensions"`
}

func DefaultTransferOperation() *TransferOperation {

	return &TransferOperation{
		DefaultAsset(),
		"1.2.0",
		"1.2.0",
		"",
		"",
		DefaultAsset(),
		DefaultMemo(),
		make([]interface{}, 0),
	}
}

func Str2Time(str string) int64 {

	str += "Z"
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return t.Unix()

}

func Time2Str(t int64) string {

	l_time := time.Unix(t, 0).UTC()
	timestr := l_time.Format(time.RFC3339)

	timestr = timestr[:len(timestr)-1]

	return timestr
}

func GetSignature(addrKey *hdkeychain.ExtendedKey, hash []byte) ([]byte, error) {

	ecPrivkey, err := addrKey.ECPrivKey()
	if err != nil {
		return nil, fmt.Errorf("in GetSignature function, get ecprivkey failed: %v", err)
	}

	ecPrivkeyByte := ecPrivkey.Serialize()
	fmt.Println("the uncompressed pubkey is: ", hex.EncodeToString(ecPrivkey.PubKey().SerializeUncompressed()))
	fmt.Println("the compressed pubkey is: ", hex.EncodeToString(ecPrivkey.PubKey().SerializeCompressed()))

	/*for {
			//sig_byte, err := hxsecp256k1.Sign(hash, ecPrivkeyByte)
			sig, err := ecPrivkey.Sign(hash)
			if err != nil {
				return nil, fmt.Errorf("in GetSignature function, get sig failed: %v", err)
			}

			sig_byte := sig.R.Bytes()
			sig_byte = append(sig_byte, sig.S.Bytes()...)
			fmt.Println(len(sig_byte))

			curve := btcec.S256()
			res := make([]byte, 1)
			for i := 0; i < (curve.H+1)*2; i++ {
				pk, err := btcec.RecoverKeyFromSignature(curve, sig, hash, i, true)
				if err != nil && pk.IsEqual(ecPrivkey.PubKey()) {
					res[0] = byte(27 + 4 + i)
					break
				} else {
					fmt.Println("get pub error!")
				}

			}

			res = append(res, sig_byte...)
			return res, nil



			//ecPubkeyByte, err := hxsecp256k1.RecoverPubkey(hash, sig_byte)
			//
			//if err != nil {
			//	return nil, fmt.Errorf("in GetSignature function, recover pubkey failed: %v", err)
			//
			//}
			//
			//// wrong with btc sign
			////sig_byte, err := btcec.SignCompact(btcec.S256(), ecPrivkey, hash, true)
			////if err != nil {
			////	return nil, fmt.Errorf("in GetSignature function, sign compact failed: %v", err)
			////}
			////
			////pub, ok, err := btcec.RecoverCompact(btcec.S256(), sig_byte, hash)
	        ////fmt.Println(ok)
			//
			//
			//if bytes.Compare(ecPubkeyByte, ecPrivkey.PubKey().SerializeUncompressed()) == 0 {
			//	fmt.Println(hex.EncodeToString(ecPubkeyByte))
			//	//fmt.Println(hex.EncodeToString(pub.SerializeCompressed()))
			//	fmt.Println(len(sig_byte))
			//	fmt.Println(hex.EncodeToString(sig_byte))
			//	fmt.Println("get pub key")
			//	return sig_byte, nil
			//}

		}*/

	for {
		sig, err := bts.SignCompact(hash, ecPrivkeyByte, true)
		if err != nil {
			return nil, fmt.Errorf("in GetSignature function, sign compact failed: %v", err)
		}

		pubkey_byte, err := bts.RecoverPubkey(hash, sig, true)
		if err != nil {
			return nil, fmt.Errorf("in GetSignature function, sign compact failed: %v", err)
		}
		fmt.Println("recoverd pubkey is: ", hex.EncodeToString(pubkey_byte))

		if bytes.Compare(ecPrivkey.PubKey().SerializeCompressed(), pubkey_byte) == 0 {
			return sig, nil
		}

	}
}

func BuildTransferTransaction(from, to, memo string, amount int64) {

	asset_amount := DefaultAsset()
	asset_amount.Hx_amount = amount

	asset_fee := DefaultAsset()
	asset_fee.Hx_amount = 2018

	memo_trx := DefaultMemo()
	memo_trx.Message = memo
	memo_trx.IsEmpty = false
	memo_trx.Hx_message = hex.EncodeToString(append(make([]byte, 4), []byte(memo_trx.Message)...))

	transferOp := DefaultTransferOperation()
	transferOp.Hx_fee = asset_fee
	transferOp.Hx_from_addr = from
	transferOp.Hx_to_addr = to
	transferOp.Hx_amount = asset_amount
	transferOp.Hx_memo = memo_trx

	expir_sec := time.Now().Unix() + 3600
	expir_str := Time2Str(expir_sec)
	//expir_str := "2018-09-13T10:06:40"
	//expir_sec := Str2Time(expir_str)

	transferTrx := TransferTransaction{
		47718,
		1191563744,
		expir_str,
		[][]interface{}{{0, transferOp}},
		make([]interface{}, 0),
		make([]string, 0),
		uint32(expir_sec),
		[]TransferOperation{*transferOp},
	}

	res := transferTrx.Serialize()
	fmt.Println("the expiration time is: ", transferTrx.Hx_expiration)

	seed := MnemonicToSeed("venture lazy digital aware plug hire acquire abuse chunk know gloom snow much employ glow rich exclude allow", "123")
	addrkey, _ := GetAddressKey(seed, 0, 0)
	addr, _ := GetAddress(seed, 0, 0)
	fmt.Println("addr is: ", addr)
	wif, _ := ExportWif(seed, 0, 0)
	fmt.Println("wif is: ", wif)

	chainid_byte, _ := hex.DecodeString(chain_id)
	toSign := sha256.Sum256(append(chainid_byte, res...))

	sig, err := GetSignature(addrkey, toSign[:])
	if err != nil {
		fmt.Println(err)
	}

	transferTrx.Hx_signatures = append(transferTrx.Hx_signatures, hex.EncodeToString(sig))
	fmt.Println("found canonical signature")
	fmt.Println(hex.EncodeToString(sig))

	b, err := json.Marshal(transferTrx)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
