package icmp

import (
	"encoding/binary"
	"fmt"

	"github.com/MoritzMy/NetMap/proto/ip"
)

// NOTE: Time Exceeded was implemented to build a custom traceroute
// Unfortunately, even L3 Switches don't decrement the TTL, therefore traceroute is semi useless in
// most local subnets

type TimeExceededPacket struct {
	ICMPHeader
	Unused uint32
	QuotedPacket
}

// Unmarshal takes a byte array of a Time Exceeded ICMP IPv4Packet and fills the fields of the Object that called
// the function. The given byte array must start right after the default ICMP Headers (start of the "Unused" Field)
func (packet *TimeExceededPacket) Unmarshal(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("time exceeded packet too small to unmarshal")
	}
	packet.Unused = binary.BigEndian.Uint32(b[0:4])

	var ipPacket QuotedPacket
	if err := ip.Unmarshal(b[4:], &ipPacket); err != nil {
		return err
	}
	return nil
}

func (packet TimeExceededPacket) Marshal() ([]byte, error) {
	// TODO: Implement me
	panic("not implemented")
}

func (packet *TimeExceededPacket) GetHeaders() *ICMPHeader {
	return &packet.ICMPHeader
}

func (packet *TimeExceededPacket) SetHeaders(header ICMPHeader) {
	packet.ICMPHeader = header
}
