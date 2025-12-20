package ethernet

import "net"

type EthernetHeader struct {
	DestinationMAC net.HardwareAddr
	SourceMAC      net.HardwareAddr
	EtherType      uint16
}
