package scan

import "net"

type InterfaceAdress struct {
	Name string
	MAC  net.HardwareAddr
	IPs  []net.IP
}

func InterfaceAdresses() []InterfaceAdress {
	ifaces, _ := net.Interfaces()
	ifaceAddrs := make([]InterfaceAdress, 0)
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		var ips []net.IP
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ips = append(ips, v.IP)
			case *net.IPAddr:
				ips = append(ips, v.IP)
			}
		}

		ifaceAddrs = append(ifaceAddrs, InterfaceAdress{
			Name: iface.Name,
			MAC:  iface.HardwareAddr,
			IPs:  ips,
		})

	}

	return ifaceAddrs
}
