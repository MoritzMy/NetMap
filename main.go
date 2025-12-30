package main

import (
	"fmt"
	"net"

	"github.com/MoritzMy/NetMap/scan"
	"github.com/MoritzMy/NetMap/scan/arp_scan"
	"github.com/MoritzMy/NetMap/scan/ping"
)

func main() {
	baddrs := scan.InterfaceAdresses()
	fmt.Println(baddrs)
	arp_scan.SendARPRequest(baddrs[0], net.IPv4(10, 254, 240, 20))

	return

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}
	ping.Sweep(addrs)
}
