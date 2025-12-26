package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/MoritzMy/NetMap/proto"
	"github.com/MoritzMy/NetMap/proto/icmp"
	"github.com/MoritzMy/NetMap/proto/ip"
)

const (
	echoReplyType = 0
)

func Ping(addr net.IP) (*ip.IPv4Packet, error) {
	var identifier uint16 = 0
	var sequenceNumber uint16 = 0

	req := icmp.NewEchoICMPPacket(identifier, sequenceNumber, []byte("ARE U UP?"))
	b, err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	conn, err := net.Dial("ip4:icmp", addr.String())

	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	err = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return nil, err
	}
	write, err := conn.Write(b)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Wrote %d Bits to: %s", write, addr.String()))

	cr := make([]byte, 200)

	err = conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return nil, err
	}
	read, err := conn.Read(cr)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v is not responding", addr))
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Recieved %d Bits from: %s", read, addr.String()))

	var packet ip.IPv4Packet

	if err := proto.Unmarshal(cr, &packet); err != nil {
		return nil, err
	}

	return &packet, nil

}
