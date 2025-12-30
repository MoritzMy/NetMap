package eth

import (
	"fmt"
	"net"
	"syscall"
)

func SendEthernetFrame(frame []byte, iface string) ([]byte, error) {
	interf, err := net.InterfaceByName(iface)

	fmt.Println("SENDING FRAME", frame)

	if err != nil {
		return nil, err
	}
	ifIndex := interf.Index

	fd, err := syscall.Socket(
		syscall.AF_PACKET,
		syscall.SOCK_RAW,
		int(htons(syscall.ETH_P_ARP)))
	if err != nil {
		return nil, err
	}
	defer syscall.Close(fd)

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ARP),
		Ifindex:  ifIndex,
	}

	if err := syscall.Bind(fd, &addr); err != nil {
		return nil, err
	}

	_, err = syscall.Write(fd, frame)

	if err != nil {
		return nil, err
	}

	response, err := ReadARPResponse(fd)
	if err != nil {
		return nil, err
	}

	return response, err
}

func htons(v uint16) uint16 {
	return (v << 8) | (v >> 8)
}
