package main

import (
	"fmt"

	"github.com/MoritzMy/NetMap/sweep/ping/icmp"
)

func main() {
	req := icmp.NewEchoICMPPacket(0, 0, []byte("Hello World! :)"))
	res, err := icmp.Marshal(&req)

	if err != nil {
		panic(fmt.Errorf("marshal echo request: %v", err))
	}

	var u icmp.EchoICMPPacket
	if err := icmp.Unmarshal(res, &u); err != nil {
		panic(err)
	}

	fmt.Println(res, "\n", u)
}
