package icmp

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	packet := NewEchoICMPPacket(0, 0, []byte("TEST"))
	cp := packet.Clone()
	b, err := Marshal(&packet)

	if err != nil {
		t.Fatal(err)
	}

	var u EchoICMPPacket

	err = Unmarshal(b, &u)
	if err != nil {
		t.Fatal(err)
	}

	if !u.Equal(cp) {
		t.Fatalf("unmarshal mismatch:\n%s\n%s", u.String(), cp.String())
	}

	if u.Checksum == cp.Checksum {
		t.Fatalf("checksum is equal after marshaling")
	}

	fmt.Println(u.String(), "\n", cp.String())

	return
}
