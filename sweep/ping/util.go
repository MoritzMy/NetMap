package ping

import (
	"encoding/binary"
	"fmt"
	"net"
)

func ValidIpsInNetwork(addr *net.IPNet) []net.IP {
	ip := addr.IP
	subnet, size := addr.Mask.Size()

	bytes := []byte(ip)

	if ip.To4() == nil {
		return nil
	}

	ipv4Bytes := bytes[len(bytes)-4:]

	fmt.Println(ipv4Bytes)

	fmt.Println(ip, subnet, size)

	if isNetworkIP(ip, subnet) || isBroadcastIP(ip, subnet) {

	}

	return nil
}

func isNetworkIP(ip net.IP, prefixLen int) bool {
	hostBits := 32 - prefixLen

	ipNumeric := binary.BigEndian.Uint32(ip)

	if ipNumeric == ipNumeric>>hostBits<<hostBits {
		return true
	}

	return false
}

func isBroadcastIP(ip net.IP, prefixLen int) bool {
	ipNumeric := binary.BigEndian.Uint32(ip)
	hostBits := 32 - prefixLen

	hostMask := uint32((1 << hostBits) - 1)

	return ipNumeric&hostMask == 0 || ipNumeric&hostMask == hostMask
}
