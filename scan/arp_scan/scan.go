package arp_scan

import (
	"fmt"
	"net"

	"github.com/MoritzMy/NetMap/proto"
	"github.com/MoritzMy/NetMap/proto/arp"
	eth "github.com/MoritzMy/NetMap/proto/ethernet"
	"github.com/MoritzMy/NetMap/proto/ip"
)

func SendARPRequest(iface net.Interface, targetIP net.IP) {
	fmt.Println("ARP Scan...")
	fmt.Println(iface, targetIP)
	fmt.Println(iface.Addrs())
	addrs, _ := iface.Addrs()

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipNet.IP.To4() == nil {
			continue
		}

		req := arp.NewARPRequest(iface.HardwareAddr, ipNet.IP, targetIP)
		b, err := proto.Marshal(&req)
		fmt.Println("RAW HEADER: ", req.EthernetHeader, "RAW ARP PACKET", req)
		if err != nil {
			panic(err)
		}
		res, err := eth.SendEthernetFrame(b, iface.Name)
		if err != nil {
			return
		}

		var hdr eth.EthernetHeader
		var pac arp.ARPRequest
		pac.EthernetHeader = &hdr

		if err := proto.Unmarshal(res, &pac); err != nil {
			panic(err)
		}
		fmt.Println("RECEIVED ARP RESPONSE:")
		fmt.Println(pac)
	}
}

func ScanNetwork(iface net.Interface) error {
	addrs, err := iface.Addrs()
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		// IPv4 only
		if ipNet.IP.To4() == nil {
			continue
		}

		for _, ip := range ip.ValidIpsInNetwork(ipNet) {
			fmt.Println("Scanning IP:", ip)
			SendARPRequest(iface, ip)
		}
	}

	return nil
}
