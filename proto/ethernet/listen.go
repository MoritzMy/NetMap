package eth

import "syscall"

// ReadARPResponse listens on the Socket stated in the fd and returns first found Responses that matches the ARP Response
// format
func ReadARPResponse(fd int) ([]byte, error) {
	buf := make([]byte, 128)

	for {
		n, _, err := syscall.Recvfrom(fd, buf, 0)

		if err != nil {
			return nil, err
		}

		if n < 42 {
			continue
		}

		if buf[12] == 0x08 && buf[13] == 0x06 && buf[21] == 2 {
			return buf[:n], nil
		}
	}
}
