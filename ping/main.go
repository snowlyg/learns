package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	pinger.Debug = true
	pinger.Interval = time.Duration(200 * time.Millisecond)
	pinger.Timeout = time.Duration(1000 * time.Millisecond)
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Printf("icmp %+v \n", stats)
}
