package eth

import (
	"net"
	"syscall"
)

func SendEthernetFrame(frame []byte, iface string) ([]byte, error) {
	interf, err := net.InterfaceByName(iface)

	if err != nil {
		return nil, err
	}

	fd, err := CreateSocket(interf)

	if err != nil {
		return nil, err
	}

	_, err = syscall.Write(fd, frame)

	if err != nil {
		return nil, err
	}

	response, err := ReadARPResponse(fd, net.IP(frame[38:42]))
	if err != nil {
		return nil, err
	}

	return response, err
}

func CreateSocket(interf *net.Interface) (int, error) {
	ifIndex := interf.Index

	fd, err := syscall.Socket(
		syscall.AF_PACKET,
		syscall.SOCK_RAW,
		int(htons(syscall.ETH_P_ARP)))
	if err != nil {
		return 0, err
	}
	defer syscall.Close(fd)

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ARP),
		Ifindex:  ifIndex,
	}

	if err := syscall.Bind(fd, &addr); err != nil {
		return 0, err
	}

	return fd, nil
}

func htons(v uint16) uint16 {
	return (v << 8) | (v >> 8)
}
