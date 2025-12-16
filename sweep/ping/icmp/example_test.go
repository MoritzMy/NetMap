package icmp_test

import (
	"fmt"
	"testing"

	"github.com/MoritzMy/NetMap/sweep/ping/icmp"
)

func TestMarshal(t *testing.T) {
	packet := icmp.NewEchoICMPPacket(0, 0, []byte("ur mom is a gae"))
	b, err := icmp.Marshal(&packet)

	if err != nil {
		panic(err)
	}

	var u icmp.EchoICMPPacket

	if err := icmp.Unmarshal(b, &u); err != nil {
		panic(err)
	}

	fmt.Println(packet, "\n", u)
}
