package hx

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/HcashOrg/hcd/chaincfg/chainhash"
	"github.com/HcashOrg/hcd/hcec/secp256k1"
	"github.com/HcashOrg/hcd/hcutil"
	"github.com/HcashOrg/hcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/ebfe/keccak"
)

const (
	CoinHC  = "HC"
	CoinBTC = "BTC"
	CoinLTC = "LTC"
	CoinETH = "eth"
)

var (
	ethSigSuffix0     = "25"
	ethSigSuffix1     = "26"
	ethSigSuffixByte0 = byte(0x25)
	ethSigSuffixByte1 = byte(0x26)
)

func SetTestnetEthSig() {
	ethSigSuffix0 = "1b"
	ethSigSuffix1 = "1c"

	ethSigSuffixByte0 = byte(0x1b)
	ethSigSuffixByte1 = byte(0x1c)
}

// SignAddress sign address to bind  to hx chain
func SignAddress(wif, address, coin string) (string, error) {
	switch coin {
	case CoinHC:
		return hcSignAddress(wif, address)

	case CoinBTC:
		return btcSignAddress(wif, address)

	case CoinLTC:
		return btcSignAddress(wif, address)

	case CoinETH:
		return ethSignAddress2(wif, address)
	}

	return "", fmt.Errorf("SignAddress: invalid coin: %s", coin)
}

func hcSignAddress(wif, addr string) (sig string, err error) {
	w, err := hcutil.DecodeWIF(wif)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, "Hc Signed Message:\n")
	wire.WriteVarString(&buf, 0, addr)
	messageHash := chainhash.HashB(buf.Bytes())

	pkCast, ok := w.PrivKey.(*secp256k1.PrivateKey)
	if !ok {
		fmt.Printf("Unable to create secp256k1.PrivateKey" +
			"from chainec.PrivateKey")
		return
	}
	res, err := secp256k1.SignCompact(secp256k1.S256(), pkCast, messageHash, true)

	return base64.StdEncoding.EncodeToString(res), nil
}

// fast hash
func Keccak256(data ...[]byte) []byte {
	h := keccak.New256()
	for _, b := range data {
		h.Write(b)
	}
	r := h.Sum(nil)

	return r
}

// eth 签名
func Sign2(wif string, msg []byte) (sig []byte, err error) {
	buf, err := hex.DecodeString(wif)
	if err != nil {
		fmt.Println("decode wif failed: ", err)
		return
	}

	key2, _ := secp256k1.PrivKeyFromBytes(secp256k1.S256(), buf)

	s2, err := secp256k1.SignCompact(secp256k1.S256(), key2, []byte(msg), false)
	return s2, err
}

// use bts sign
func ethSignAddress2(wif, addr string) (data string, err error) {
	baddr, _ := hex.DecodeString(addr)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(baddr))

	h := Keccak256(append([]byte(msg), baddr...))

	sig, err := Sign2(wif, h)
	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])

	if v == byte(0) {
		sig[64] = ethSigSuffixByte0
		// res = res[0:len(res)-2] + "1b"
	} else {
		sig[64] = ethSigSuffixByte1
		// res = res[0:len(res)-2] + "1c"
	}

	return "0x" + hex.EncodeToString(sig), nil
}

/*
func ethSignAddress(wif, addr string) (sig string, err error) {
	baddr, _ := hex.DecodeString(addr)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(baddr))

	h := crypto.Keccak256(append([]byte(msg), baddr...))
	// fmt.Println("msg addr:", hex.EncodeToString(h))

	buf, err := hex.DecodeString(wif)
	if err != nil {
		fmt.Println("decode wif failed: ", err)
		return
	}

	key, err := crypto.ToECDSA(buf)

	data, err := crypto.Sign(h, key)

	if err != nil {
		fmt.Printf("sign eth failed: %v", err)
		return
	}
	// fmt.Println("signed:", hex.EncodeToString(data))

	// TODO: 测试链和正式链的结尾不同
	// 测试链 00 -> 1b   01 -> 1c
	// 正式链 00 -> 25   01 -> 26
	res := hex.EncodeToString(data)
	suffix := res[len(res)-2 : len(res)]
	if suffix == "00" {
		res = res[0:len(res)-2] + "1b"
	} else if suffix == "01" {
		res = res[0:len(res)-2] + "1c"
	} else {
		return "", fmt.Errorf("invalid signature suffix: %v", suffix)
	}

	return "0x" + res, nil
}
*/

func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func btcSignAddress(wif, addr string) (sig string, err error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, "Bitcoin Signed Message:\n")
	wire.WriteVarString(&buf, 0, addr)

	messageHash := DoubleHashB(buf.Bytes())

	pkCast := secp256k1.PrivateKey(*w.PrivKey)

	res, err := secp256k1.SignCompact(secp256k1.S256(), &pkCast, messageHash, true)

	return base64.StdEncoding.EncodeToString(res), nil
}

func ltcSignAddress(wif, addr string) (sig string, err error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, "Litecoin Signed Message:\n")
	wire.WriteVarString(&buf, 0, addr)

	messageHash := DoubleHashB(buf.Bytes())

	pkCast := secp256k1.PrivateKey(*w.PrivKey)
	res, err := secp256k1.SignCompact(secp256k1.S256(), &pkCast, messageHash, true)

	return base64.StdEncoding.EncodeToString(res), nil
}

func btsSign(wif string, data []byte) (res []byte, err error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return
	}

	pkCast := secp256k1.PrivateKey(*w.PrivKey)
	fmt.Println("ecPrivkey", *pkCast.D)
	res, err = secp256k1.SignCompact(secp256k1.S256(), &pkCast, data, true)
	return
}
