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

type EchoRequest struct {
	messageType    uint8
	code           uint8
	checksum       uint16
	identifier     uint16
	sequenceNumber uint16
	payload        []byte
}

func CreateEchoRequest(identifier uint16, sequenceNumber uint16, payload []byte) EchoRequest {
	return EchoRequest{
		messageType:    echoType,
		code:           echoCode,
		identifier:     identifier,
		sequenceNumber: sequenceNumber,
		payload:        payload,
	}
}

func Marshal(request EchoRequest) ([]byte, error) {
	if len(request.payload) > maxPayload {
		return nil, fmt.Errorf("marshal icmp request: payload size %d exceeds limit of %d Bytes", len(request.payload), maxPayload)
	}

	b := make([]byte, 0, 8+len(request.payload))
	b = append(b, request.messageType, request.code)
	b = binary.BigEndian.AppendUint16(b, checksumPlaceholder)
	b = binary.BigEndian.AppendUint16(b, request.identifier)
	b = binary.BigEndian.AppendUint16(b, request.sequenceNumber)
	b = append(b, request.payload...)

	return b, nil
}

func computeChecksum(request []byte) uint16 {
	return 0
}
