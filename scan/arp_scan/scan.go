package arp_scan

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/MoritzMy/NetMap/proto"
	"github.com/MoritzMy/NetMap/proto/arp"
	eth "github.com/MoritzMy/NetMap/proto/ethernet"
	"github.com/MoritzMy/NetMap/proto/ip"
)

func SendARPRequest(iface net.Interface, targetIP net.IP) bool {
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
		if err != nil {
			log.Println("error occurred while marshalling ARP request:", err)
			return false
		}
		res, err := eth.SendEthernetFrame(b, iface.Name)
		if err != nil {
			log.Println("error occurred while sending ARP request:", err)
			return false
		}

		var hdr eth.EthernetHeader
		var pac arp.ARPRequest
		pac.EthernetHeader = &hdr

		if err := proto.Unmarshal(res, &pac); err != nil {
			log.Println("error occurred while unmarshalling ARP response:", err)
			return false
		}
		fmt.Println("RECEIVED ARP RESPONSE:")
		fmt.Println(pac)
	}
	return true

}

func ScanNetwork(iface net.Interface) error {
	addrs, err := iface.Addrs()
	if err != nil {
		return err
	}

	respondedIps := make([]net.IP, 0)

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		// IPv4 only
		if ipNet.IP.To4() == nil {
			continue
		}

		var wg sync.WaitGroup

		for _, ip := range ip.ValidIpsInNetwork(ipNet) {
			fmt.Println("Scanning IP:", ip)
			wg.Go(func() {
				if ok := SendARPRequest(iface, ip); !ok {
					log.Println("Failed to send ARP request to", ip)
				} else {
					respondedIps = append(respondedIps, ip)
				}
			})
		}

		wg.Wait()
	}

	fmt.Println(respondedIps)

	return nil

}
