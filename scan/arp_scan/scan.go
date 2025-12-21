package arp_scan

import (
	"net"

	"github.com/MoritzMy/NetMap/proto/arp"
	"github.com/MoritzMy/NetMap/scan"
)

func ARPScan(adress scan.InterfaceAdress) {
	arp.NewARPRequest(adress.MAC, adress.IPs[0], net.IP{0x0a, 0xfe, 0xf0, 0x13})
}
