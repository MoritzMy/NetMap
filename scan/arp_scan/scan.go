package arp_scan

import (
	"fmt"
	"log"
	"net"

	"github.com/MoritzMy/NetMap/proto"
	"github.com/MoritzMy/NetMap/proto/arp"
	eth "github.com/MoritzMy/NetMap/proto/ethernet"
	"github.com/MoritzMy/NetMap/scan"
)

func SendARPRequest(address scan.InterfaceAddress, targetIP net.IP) {
	fmt.Println("ARP Scan...")
	fmt.Println(address, targetIP)
	fmt.Println(address.IPs[0])
	req := arp.NewARPRequest(address.MAC, address.IPs[0], net.IP{10, 254, 240, 20})
	b, err := proto.Marshal(&req)
	fmt.Println("RAW HEADER: ", req.EthernetHeader, "RAW ARP PACKET", req)
	if err != nil {
		panic(err)
	}
	res, err := eth.SendEthernetFrame(b, "enp0s31f6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent")
	fmt.Println("Response: ", string(res))
}
