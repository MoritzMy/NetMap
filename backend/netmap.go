package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"

	arp_scan2 "github.com/MoritzMy/NetMap/backend/cmd/arp_scan"
	"github.com/MoritzMy/NetMap/backend/cmd/ping"
	graphing2 "github.com/MoritzMy/NetMap/backend/internal/graphing"
)

func main() {
	arp := flag.Bool("arp-scan", false, "Run ARP Discovery Scan")
	icmp := flag.Bool("ping-sweep", false, "Run ICMP Sweep")
	dot_file := flag.String("dot-file", "", "Output the resulting graph to a DOT file")

	flag.Parse()

	graph := graphing2.NewGraph()

	if *arp {
		runARPScan(graph)
	}

	if *icmp {
		runICMPSweep(graph)
	}

	if !*arp && !*icmp {
		fmt.Println("Please specify a scan type. Use -h for help.")
	}

	for node := range graph.Nodes {
		graph.GetOrCreateNode(node).EnrichNode() // Enrich nodes with additional information
	}

	graph.LinkNetworkToGateway()

	json, err := graph.MarshalJSON()

	if err != nil {
		fmt.Printf("Could not create json file", err)
	}

	fmt.Printf(string(json))

	wd, err := os.Getwd()

	outPath := filepath.Join(wd, "graph.json")

	err = os.WriteFile(outPath, json, 0644)

	if err != nil {
		fmt.Printf("Could not Write to File")
	}

	if *dot_file != "" {
		err := graph.ExportToDOT(*dot_file)
		if err != nil {
			fmt.Println("Error exporting graph to DOT file:", err)
		} else {
			fmt.Printf("Graph exported to %s\n", *dot_file)
		}
	}

	fmt.Println(graph)
}

func runARPScan(graph *graphing2.Graph) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}

	in := make(chan arp_scan2.ARPEvent)

	go func() {
		for ev := range in {
			fmt.Printf("Discovered device - IP: %s, MAC: %s, Network: %s, Source: %s\n", ev.IP, ev.MAC, ev.Network, ev.Source)
			node := graph.GetOrCreateNode("ip:" + ev.IP.String())
			node.MAC = ev.MAC
			node.IP = ev.IP
			node.Protocols["arp"] = true
			netNode := graph.GetOrCreateNode("net:" + ev.Network.String())
			netNode.Type = graphing2.NodeNetwork
			graph.AddEdge(node.ID, netNode.ID, graphing2.EdgeMemberOf)
		}
	}()

	for _, iface := range ifaces {
		fmt.Printf("Starting ARP Scan on interface %s\n", iface.Name)
		err := arp_scan2.Scan(iface, in)
		if err != nil {
			fmt.Printf("Error scanning network on interface %s: %v\n", iface.Name, err)
			continue
		}
	}

}

func runICMPSweep(graph *graphing2.Graph) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}
	in := make(chan net.IP)

	go func() {
		for ip := range in {
			fmt.Printf("Discovered host - IP: %s\n", ip)
			node := graph.GetOrCreateNode("ip:" + ip.String())
			node.Protocols["icmp"] = true
		}
	}()

	for _, iface := range ifaces {
		fmt.Printf("Starting ICMP Sweep on interface %s\n", iface.Name)
		if err := ping.Sweep(iface, in); err != nil {
			fmt.Printf("Error during ICMP Sweep on interface %s: %v\n", iface.Name, err)
		}
	}
}
