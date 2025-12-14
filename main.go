package main

import (
	"fmt"
	"log"
	"net"

	ping "github.com/MoritzMy/NetMap/sweep/ping"
)

func main() {
	ifaces, _ := net.Interfaces()

	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)

			if !ok {
				continue
			}

			// Avoid the Loopback IP since it's not relevant for scan and any non IPv4 IPs
			if ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}
			ip4Addr, ip4Net, err := net.ParseCIDR(ipNet.String())

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(ip4Addr, ip4Net)

			ping.Ping(ip4Net)

		}
	}
}
