package main

import (
	"fmt"

	"github.com/wens07/eth_lib"
)

func main() {

	balance := eth_lib.ETH_getBalance_by_block("0xfd5edcd98d3d009915e0f203610e5dd747e6d005", 4730666)
	fmt.Println(balance)
	fmt.Println(balance[:len(balance)-18])

}
