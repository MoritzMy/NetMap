package arp

import "net"

type ARPPayload struct {
	HTYPE     uint16
	PTYPE     uint16
	HLEN      uint8
	PLEN      uint8
	OPER      uint16
	SourceMAC net.HardwareAddr
	SourceIP  net.IP
	TargetMAC net.HardwareAddr
	TargetIP  net.IP
}
