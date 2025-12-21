package main

import (
	"net"

	"github.com/MoritzMy/NetMap/scan/ping"
)

func main() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}
	ping.Sweep(addrs)
}
