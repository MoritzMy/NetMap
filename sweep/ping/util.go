package ping

import (
	"fmt"
	"net"
)

func ValidIpsInNetwork(addr *net.IPNet) bool {
	ip := addr.IP
	subnet, size := addr.Mask.Size()

	bytes := []byte(ip)

	if ip.To4() == nil {
		return false
	}

	ipv4Bytes := bytes[len(bytes)-5:]

	fmt.Println(ipv4Bytes)

	fmt.Println(ip, subnet, size)

	return true
}
