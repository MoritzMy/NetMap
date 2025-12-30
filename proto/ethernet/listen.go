package eth

import (
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
	"time"
)

// ReadARPResponse listens on the Socket stated in the fd and returns first found Responses that matches the ARP Response
// format
func ReadARPResponse(fd int, targetIP net.IP) ([]byte, error) {
	buf := make([]byte, 128)

	timer := time.NewTimer(3 * time.Second)

	for {
		n, _, err := syscall.Recvfrom(fd, buf, 0)

		if err != nil {
			return nil, err
		}

		if n < 42 {
			continue
		}

		if buf[12] == 0x08 && buf[13] == 0x06 && binary.BigEndian.Uint16(buf[20:22]) == 2 && net.IP(buf[28:32]).Equal(targetIP) {
			return buf[:n], nil
		}

		<-timer.C
		return nil, fmt.Errorf(syscall.ETIMEDOUT.Error(), "to ", targetIP)
	}
}
