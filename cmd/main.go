package main

import (
	"fmt"
	"time"

	"mkid"
)

func main() {
	// res, lasttime, seq := mkid(time.Now().Unix(), 0)
	var res uint64
	lasttime, seq := time.Now().Unix(), 0
	for i := 0; i < 10000; i++ {
		res, lasttime, seq = mkid.Mkid(lasttime, seq)
		fmt.Printf("%d,", seq)
		fmt.Printf("%x\n", res)
		// time.Sleep(time.Millisecond * 20)
	}
}
