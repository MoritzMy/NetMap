package ping

import (
	"encoding/binary"
	"fmt"
)

const (
	echoCode            = 0
	echoType            = 8
	checksumPlaceholder = 0
	maxPayload          = 56
)

type Marshaler interface {
	Marshal() ([]byte, error)
}

type ICMPPacket struct {
	MessageType uint8
	Code        uint8
	Payload        []byte

}

type EchoRequest struct {
	ICMPPacket
	Identifier     uint16
	SequenceNumber uint16
}

func CreateEchoRequest(identifier uint16, sequenceNumber uint16, payload []byte) EchoRequest {
	return EchoRequest{
		ICMPPacket: ICMPPacket{
			MessageType: echoType,
			Code:        echoCode,
			Payload:     payload,

		},
		Identifier:     identifier,
		SequenceNumber: sequenceNumber,
	}
}

func (packet ICMPPacket) Marshal() ([]byte, error) {
	if len(packet.Payload) > maxPayload {
		return nil, fmt.Errorf("payload too large")
	}
	
	b:= make([]byte, 0, 4 + len(packet.Payload))
	b = append(b, packet.MessageType, packet.Code)
	b = binary.BigEndian.AppendUint16(b, checksumPlaceholder)
	b = append(b, packet.Payload...)
	
	cs := computeChecksum(b)

	binary.BigEndian.PutUint16(b[2:4], cs)

	return b, nil
}

func (request EchoRequest) Marshal() ([]byte, error) {
	if len(request.Payload) > maxPayload {
		return nil, fmt.Errorf("marshal icmp request: payload size %d exceeds limit of %d Bytes", len(request.Payload), maxPayload)
	}

	b := make([]byte, 0, 8+len(request.Payload))
	b = append(b, request.MessageType, request.Code)
	b = binary.BigEndian.AppendUint16(b, checksumPlaceholder)
	b = binary.BigEndian.AppendUint16(b, request.Identifier)
	b = binary.BigEndian.AppendUint16(b, request.SequenceNumber)
	b = append(b, request.Payload...)

	cs := computeChecksum(b)

	binary.BigEndian.PutUint16(b[2:4], cs)

	return b, nil
}

// computeChecksum computes the checksum of a package, by splitting it up into 16 Bit words,
// adding those words together and performing an end around carry until the sum is also a 16 Bit word.
// In the Case of ICMP, while the Checksum is not computed, a Placeholder should be used of which the 16 Bit
// word value is 0
func computeChecksum(request []byte) uint16 {
	sum := uint32(0)

	// Turn the bytes into 16 Bit Words and add them up
	for i := 0; i+1 < len(request); i += 2 {
		sum += (uint32(request[i]) << 8) + uint32(request[i+1])
	}

	if len(request)%2 != 0 {
		sum += uint32(request[len(request)-1]) << 8
	}

	// sum needs to be a valid uint16, otherwise an end around carry is performed
	for sum>>16 != 0 {
		sum = uint32(uint16(sum)) + sum>>16
	}

	// One's complement
	var checksum = ^uint16(sum)

	return checksum
}
