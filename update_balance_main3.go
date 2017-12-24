package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/wens07/eth_lib"
)

func main() {

	runtime.GOMAXPROCS(4)

	var start_row uint64 = 8055000
	var end_row uint64 = 8255000
	var step uint64 = 1000

	start := time.Now()

	for start_row < end_row {

		var wg sync.WaitGroup
		wg.Add(4)

		go eth_lib.Multigoroutine_update_balance(start_row, step, &wg)
		go eth_lib.Multigoroutine_update_balance(start_row+step, step, &wg)
		go eth_lib.Multigoroutine_update_balance(start_row+2*step, step, &wg)
		go eth_lib.Multigoroutine_update_balance(start_row+3*step, step, &wg)

		wg.Wait()

		start_row = start_row + 4*step

		fmt.Println("current row: ", start_row)
	}

	duration := time.Since(start)

	fmt.Println("executation time: ", duration)

}