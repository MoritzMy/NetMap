package main

import (
	"fmt"
	"net"

	"github.com/MoritzMy/NetMap/scan/arp_scan"
	"github.com/MoritzMy/NetMap/scan/ping"
)

func main() {
	baddrs, err := net.Interfaces()
	fmt.Println(baddrs)
	arp_scan.ScanNetwork(baddrs[1])

	return

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}
}
