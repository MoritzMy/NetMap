package ping

import (
	"encoding/binary"
	"net"
)

func ValidIpsInNetwork(addr *net.IPNet) []net.IP {
	var hosts []net.IP

	baseAddr := addr.IP.Mask(addr.Mask)

	ip := addr.IP
	subnet, size := addr.Mask.Size()

	ip4 := ip.To4()

	if ip4 == nil {
		return nil
	}

	currAddr := baseAddr

	for i := 0; i < 2^(size-subnet); i++ {
		IncrementIP(currAddr)
		if isNetworkIP(ip, subnet) || isBroadcastIP(ip, subnet) {
			continue
		}
		hosts = append(hosts, ip)
	}

	return hosts
}

func IncrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] != 0 {
			break
		}
	}
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
