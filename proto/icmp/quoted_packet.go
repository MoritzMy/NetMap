package icmp

import "github.com/MoritzMy/NetMap/proto/ip"

type QuotedPacket struct {
	Header  ip.Header
	Payload [8]byte
}

func (q QuotedPacket) Marshal() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (q QuotedPacket) Unmarshal(bytes []byte) error {
	q.Payload = [8]byte(bytes[q.HeaderSize() : q.HeaderSize()+8])
	return nil
}

func (q QuotedPacket) HeaderSize() int {
	return int(q.Header.VersionIHL.Size())
}

func (q QuotedPacket) Size() int {
	return int(q.Header.VersionIHL.Size()) + 8
}

func (q QuotedPacket) GetHeaders() *ip.Header {
	return &q.Header
}

func (q QuotedPacket) SetHeaders(header ip.Header) {
	q.Header = header
}
